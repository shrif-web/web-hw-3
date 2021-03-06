package main

import (
	"context"
	"database/sql"
	"fmt"
	pb "hw3/BackEnd/cacheproto"
	. "hw3/BackEnd/the_cache"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"google.golang.org/grpc"
)

type user struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	UserId        int    `bun:"user_id,pk,autoincrement"`
	UserName      string `bun:"user_name,notnull"`
	Name          string `bun:"name,notnull"`
	Password      string `bun:"password,notnull"`
}

//type Note struct {
//	bun.BaseModel `bun:"table:notes,alias:u"`
//	NoteId        int    `bun:"note_id,pk,autoincrement"`
//	Note          string `bun:"Note,notnull"`
//	NoteTitle     string `bun:"title,notnull"`
//	NoteType      string `bun:"type,notnull"`
//	AuthorId      int    `bun:"author_id"`
//}

func (u user) String() string {
	return fmt.Sprintf("User<%d %s %s>", u.UserId, u.UserName, u.Password)
}

const (
	port = ":50051"
)

type CacheManagementServer struct {
	pb.UnimplementedCacheManagementServer
}

var cache TheCache
var db *bun.DB

func toMyNote(notes []Note) []*pb.Note {
	var pbNotes []*pb.Note
	pbNotes = make([]*pb.Note, len(notes))
	for i := 0; i < len(notes); i++ {
		pbNotes[i] = &pb.Note{
			Text:  notes[i].Note,
			Title: notes[i].NoteTitle,
			Id:    strconv.Itoa(notes[i].NoteId),
			Type:  notes[i].NoteType,
		}
	}
	return pbNotes
}
func probNotesToNotes(notes []*pb.Note) []Note {
	var pbNotes []Note
	pbNotes = make([]Note, len(notes))
	for i := 0; i < len(notes); i++ {
		nId, _ := strconv.Atoi(notes[i].Id)
		pbNotes[i] = Note{
			Note:      notes[i].Text,
			NoteTitle: notes[i].Title,
			NoteType:  notes[i].Type,
			NoteId:    nId,
		}
	}
	return pbNotes
}
func findNodeById(notes []Note, nId int) *Note {
	for _, note := range notes {
		if note.NoteId == nId {
			return &note
		}
	}
	return nil
}
func (s *CacheManagementServer) CacheLoginRPC(in *pb.CacheLoginRequest, a pb.CacheManagement_CacheLoginRPCServer) error {
	//todo handle request
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Minute)
	var res *pb.CacheLoginResponse
	res = &pb.CacheLoginResponse{}
	//todo The cache
	//check cache
	//back to DB if not exist
	switch in.RequestType {
	case Login:
		node := cache.GetUserKey(in.User, in.Pass)
		if node == nil {
			res.MissCache = true
			{
				node = new(Node)
				userObj := &user{}
				err := db.NewSelect().Model(userObj).Where("user_name = ? AND password = ?", in.User, in.Pass).Scan(ctx)
				if err != nil {
					fmt.Println(err)
					res.WrongPass = true
				} else {
					res.UserId = strconv.Itoa(userObj.UserId)
					res.Exist = true
					var notes []Note
					err = db.NewSelect().Model(&notes).Where("author_id = ?", res.UserId).Scan(ctx)
					res.Notes = toMyNote(notes)
					res.UserName = userObj.UserName
					res.Name = userObj.Name
					if res.UserId != "0" {
						node.UserId = userObj.UserId
						node.UserName = userObj.UserName
						node.Name = userObj.Name
						node.Password = userObj.Password
						node.Notes = notes
						cache.SetKey(node)
					}
				}
			}
		} else {
			res.Exist = true
			res.WrongPass = false
			res.Notes = toMyNote(node.Notes)
			res.UserName = node.UserName
			res.UserId = strconv.Itoa(node.UserId)
			res.Name = node.Name
		}
	case SignUp:
		userObj := &user{}
		err := db.NewSelect().Model(userObj).Where("user_name = ?", in.User).Scan(ctx)
		if err == nil {
			res.Exist = true
		} else if err == sql.ErrNoRows {
			res.Exist = false
			userObj = &user{
				BaseModel: bun.BaseModel{},
				UserName:  in.User,
				Password:  in.Pass,
				Name:      in.Name,
			}
			_, err := db.NewInsert().Model(userObj).Exec(ctx)
			if err == nil {
				node := new(Node)
				res.UserId = strconv.FormatInt(int64(userObj.UserId), 10)
				res.Name = userObj.Name
				res.UserName = userObj.UserName
				if res.UserId != "0" {
					node.Name = userObj.Name
					node.UserId = userObj.UserId
					node.UserName = userObj.UserName
					node.Password = userObj.Password
					fmt.Println("new user cache node made")
					cache.SetKey(node)
				}
			}
		}
	}
	//write to cache if not exist
	err := a.Send(res)
	if err != nil {
		return err
	}
	return nil
}
func (s *CacheManagementServer) CacheNoteRPC(in *pb.CacheNoteRequest, a pb.CacheManagement_CacheNoteRPCServer) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	var res *pb.CacheNoteResponse
	res = &pb.CacheNoteResponse{
		Note:      "",
		NoteId:    "",
		Exist:     true,
		Access:    false,
		MissCache: false,
	}
	cacheData := &CacheData{}
	switch in.RequestType {
	case Save:
		{
			aId, _ := strconv.Atoi(in.AuthorId)
			noteObj := &Note{
				BaseModel: bun.BaseModel{},
				Note:      in.Note,
				AuthorId:  aId,
				NoteTitle: in.NoteTitle,
				NoteType:  in.Type,
			}
			_, err := db.NewInsert().Model(noteObj).Returning("*").Exec(ctx)
			if err == nil {
				res.NoteId = strconv.FormatInt(int64(noteObj.NoteId), 10)
			}
			cacheData.CommandType = Save
			cacheData.UserId = aId
			cacheData.Notes = []Note{*noteObj}
			exist := cache.SetExistingKey(cacheData)
			if !exist {
				res.MissCache = true
				createCacheNode(aId, ctx)
			}
		}
	case Del:
		nId, _ := strconv.Atoi(in.NoteId)
		aId, _ := strconv.Atoi(in.AuthorId)
		noteObj := &Note{NoteId: nId, AuthorId: aId}
		var err error
		var r sql.Result
		if aId == 0 {
			// todo cache
			r, err = db.NewDelete().Model(noteObj).Where("note_id = ?", nId).Exec(ctx)
		} else {
			r, err = db.NewDelete().Model(noteObj).Where("note_id = ? AND author_id = ?", nId, aId).Exec(ctx)
		}
		if err == nil {
			rows, _ := r.RowsAffected()
			if rows > 0 {
				res.Exist = true
				res.Access = true
				if aId != 0 {
					cacheData.CommandType = Del
					cacheData.UserId = aId
					cacheData.Notes = []Note{*noteObj}
					exist := cache.SetExistingKey(cacheData)
					if !exist {
						res.MissCache = true
						createCacheNode(aId, ctx)
					}
				}
			}
		} else {
			print(err)
			res.Exist = false
			res.Access = false
		}
	case GetAll:
		aid, _ := strconv.Atoi(in.AuthorId)
		node := cache.GetKey(aid)
		if node == nil {
			res.MissCache = true
			{
				node = new(Node)
				var notes []Note
				var err error
				if aid == 0 {
					err = db.NewSelect().Model(&notes).Scan(ctx)

				} else {
					err = db.NewSelect().Model(&notes).Where("author_id = ?", in.AuthorId).Scan(ctx)
				}
				if err != nil {
					fmt.Println(err)
					res.Exist = false
				} else {
					res.Notes = toMyNote(notes)
					res.Exist = true
					res.Access = true
					if aid != 0 {
						userObj := &user{}
						err := db.NewSelect().Model(userObj).Where("user_id = ?", aid).Scan(ctx)
						if err != nil {
							// todo fosh
						}
						node.UserId, _ = strconv.Atoi(in.AuthorId)
						node.Notes = notes
						node.Name = userObj.Name
						node.UserName = userObj.UserName
						node.Password = userObj.Password
						cache.SetKey(node)
					}
				}
			}
		} else {
			res.Exist = true
			res.Access = true
			res.Notes = toMyNote(node.Notes)
		}
	case Get:
		aid, _ := strconv.Atoi(in.AuthorId)
		node := cache.GetKey(aid)
		if node == nil {
			res.MissCache = true
			{
				noteObj := &Note{}
				err := db.NewSelect().Model(noteObj).Where("note_id = ?", in.NoteId).Scan(ctx)
				if err != nil {
					fmt.Println(err)
					res.Exist = false
				} else {
					aid, _ := strconv.Atoi(in.AuthorId)
					res = &pb.CacheNoteResponse{
						Note:      noteObj.Note,
						NoteId:    strconv.Itoa(noteObj.NoteId),
						Title:     noteObj.NoteTitle,
						Type:      noteObj.NoteType,
						Exist:     true,
						Access:    in.AuthorId == strconv.Itoa(noteObj.AuthorId) || in.AuthorId == "0",
						MissCache: true,
					}
					if in.AuthorId != "0" {
						createCacheNode(aid, ctx)
					}
				}
			}
		} else {
			nId, _ := strconv.Atoi(in.NoteId)
			note := findNodeById(node.Notes, nId)
			if note != nil {
				res.Exist = true
				res.Access = true
				res.Note = note.Note
				res.Title = note.NoteTitle
				res.Type = note.NoteType
				res.NoteId = in.NoteId
			} else {
				res.Exist = false
				res.Access = false
			}
		}
	case Edit:
		aId, _ := strconv.Atoi(in.AuthorId)
		nId, _ := strconv.Atoi(in.NoteId)
		noteObj := &Note{
			BaseModel: bun.BaseModel{},
			Note:      in.Note,
			NoteTitle: in.NoteTitle,
			NoteId:    nId,
			AuthorId:  aId,
			NoteType:  in.Type,
		}
		var err error
		if aId == 0 {
			_, err = db.NewUpdate().Model(noteObj).Where("note_id = ?", nId).Exec(ctx)
		} else {
			_, err = db.NewUpdate().Model(noteObj).Where("note_id = ? AND author_id = ?", nId, aId).Exec(ctx)
		}
		if err == nil {
			res.Access = true
			res.Exist = true
			if aId != 0 {
				cacheData.CommandType = Edit
				cacheData.UserId, _ = strconv.Atoi(in.AuthorId)
				cacheData.Notes = []Note{*noteObj}
				exist := cache.SetExistingKey(cacheData)
				if !exist {
					res.MissCache = true
					createCacheNode(aId, ctx)
				}
			}
		} else {
			res.Access = false
			res.Exist = false
		}
	}
	err := a.Send(res)
	if err != nil {
		return err
	}
	return nil
}
func connectToDB() {
	// Open a PostgreSQL database.
	dsn := "postgres://postgres:test123@localhost:5432/postgres?sslmode=disable"
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	// Create a Bun db on top of it.
	db = bun.NewDB(pgdb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	//userObj := &user{}
	//err := db.NewSelect().Model(userObj).Where("user_name = ? AND password = ?", "amir", "Xamm2666").Scan(context.Context())
}
func createCacheNode(uId int, ctx context.Context) {
	userObj := &user{}
	_ = db.NewSelect().Model(userObj).Where("user_id = ?", uId).Scan(ctx)
	var notes []Note
	_ = db.NewSelect().Model(&notes).Where("author_id = ?", uId).Scan(ctx)
	node := new(Node)
	node.Notes = notes
	node.UserName = userObj.UserName
	node.Name = userObj.Name
	node.Password = userObj.Password
	node.UserId = uId
	cache.SetKey(node)
}
func main() {
	connectToDB()
	cache = InitCache()
	startGrpcServer()
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCacheManagementServer(s, &CacheManagementServer{})
	//pb.RegisterRequestCacheServer(s, &RequestCacheServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
