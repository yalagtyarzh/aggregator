import {Component, Fragment} from "react";
import Alert from "./ui-components/Alert";
import Input from "./form-components/Input";

export default class Login extends Component {
    constructor(props) {
        super(props);

        this.state = {
            userName: "",
            password: "",
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
        let value = this.target.value;
        let name = this.target.name;
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
                    <Input title={"Username"} type={"text"} name={"userName"} handleChange={this.handleChange}
                           className={this.hasError("userName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("userName") ? "test-danger" : "d-none"}
                           errorMsg={"Please enter a valid username"}/>

                    <Input title={"Password"} type={"password"} name={"password"} handleChange={this.handleChange}
                           className={this.hasError("password") ? "is-invalid" : ""}
                           errorDiv={this.hasError("password") ? "test-danger" : "d-none"}
                           errorMsg={"Please enter a password"}/>

                    <hr/>

                    <button className={"btn btn-primary"}>Login</button>
                </form>
            </Fragment>
        );
    }
}