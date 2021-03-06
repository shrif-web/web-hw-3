import React from "react";
import Joi from "joi-browser";
import Form from "./common/form";
import { login } from "../services/authService";
import { toast } from "react-toastify";

class LoginForm extends Form {
  state = {
    data: { username: "", password: "" },
    errors: {},
  };

  schema = {
    username: Joi.string().required().label("Username"),
    password: Joi.string().required().label("Password"),
  };

  doSubmit = async () => {
    // backend
    console.log("login");
    const { data } = this.state;
    const output = await login(data.username, data.password);
    console.log(output);
    if (output.data && output.data.misscache) {
      toast.warn("MissCached in the GET(all) request happened");
    }
    localStorage.setItem("jwt", output.data["jwt"]);
    localStorage.setItem("notes", JSON.stringify(output.data["notes"]));
    localStorage.setItem("name", output.data["name"]);
    window.location = "/";
    // try catch 400 username not exists
    console.log("Submitted");
  };

  render() {
    return (
      <div>
        <h1>Login</h1>
        <form onSubmit={this.handleSubmit}>
          {this.renderInput("username", "Username")}
          {this.renderInput("password", "Password", "password")}
          {this.renderButton("Login")}
        </form>
      </div>
    );
  }
}

export default LoginForm;
