import {Component, Fragment} from "react";
import Alert from "./ui-components/Alert";
import Input from "./form-components/Input";

export default class Register extends Component {
    constructor(props) {
        super(props);

        this.state = {
            firstName: "",
            lastName: "",
            userName: "",
            email: "",
            password: "",
            confirm: "",
            error: null,
            errors: [],
            alert: {
                type: "d-none",
                message: "",
            }
        }

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange = (evt) => {
        let value = evt.target.value;
        let name = evt.target.name;
        this.setState((prevState) => ({
            ...prevState,
            [name]: value,
        }));
    }

    handleSubmit = (evt) => {
        evt.preventDefault();
    }

    hasError(key) {
        return this.state.errors.indexOf(key) !== -1;
    }

    render() {
        return (
            <Fragment>
                <h2>Login</h2>
                <hr/>
                <Alert alertType={this.state.alert.type}
                       alertMessage={this.state.alert.message}/>


                <form className={"pt-3"} onSubmit={this.handleSubmit}>
                    <Input title={"First name"} type={"text"} name={"firstName"} handleChange={this.handleChange}
                           className={this.hasError("firstName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("firstName") ? "text-danger" : "d-none"}
                           errorMsg={"Please enter a valid first name"}/>

                    <Input title={"Last name"} type={"text"} name={"lastName"} handleChange={this.handleChange}
                           className={this.hasError("lastName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("lastName") ? "text-danger" : "d-none"}
                           errorMsg={"Please enter a valid last name"}/>

                    <Input title={"Username"} type={"text"} name={"userName"} handleChange={this.handleChange}
                           className={this.hasError("userName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("userName") ? "text-danger" : "d-none"}
                           errorMsg={"Please enter a valid username"}/>

                    <Input title={"Email"} type={"email"} name={"email"} handleChange={this.handleChange}
                           className={this.hasError("email") ? "is-invalid" : ""}
                           errorDiv={this.hasError("email") ? "text-danger" : "d-none"}
                           errorMsg={"Please enter a valid email"}/>

                    <Input title={"Password"} type={"password"} name={"password"} handleChange={this.handleChange}
                           className={this.hasError("password") ? "is-invalid" : ""}
                           errorDiv={this.hasError("password") ? "text-danger" : "d-none"}
                           errorMsg={"Please enter a password"}/>

                    <Input title={"Confirm password"} type={"text"} name={"confirm"} handleChange={this.handleChange}
                           className={this.hasError("confirm") ? "is-invalid" : ""}
                           errorDiv={this.hasError("confirm") ? "text-danger" : "d-none"}
                           errorMsg={"Please enter a password confirmation"}/>

                    <hr/>

                    <button className={"btn btn-primary"}>Sign in</button>
                </form>
            </Fragment>
        );
    }
}