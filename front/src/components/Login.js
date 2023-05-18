import {Component, Fragment} from "react";
import Alert from "./ui-components/Alert";
import Input from "./form-components/Input";
import jwt_decode from "jwt-decode";

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
        let value = evt.target.value;
        let name = evt.target.name;
        this.setState((prevState) => ({
            ...prevState,
            [name]: value,
        }));
    }

    handleSubmit = (evt) => {
        evt.preventDefault();

        let errors = [];

        if (this.state.userName === "" || this.state.userName.length < 3 || this.state.userName.length > 32) {
            errors.push("userName");
        }

        if (this.state.password === "" || this.state.password.length < 3 || this.state.password.length > 32) {
            errors.push("password");
        }

        this.setState({errors: errors});

        if (errors.length > 0) {
            return false;
        }

        const req = {
            userName: this.state.userName,
            password: this.state.password,
        }

        const requestOptions = {
            method: "POST",
            body: JSON.stringify(req)
        }

        fetch("http://localhost/api/v1/login", requestOptions)
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    this.setState({
                        alert: {
                            type: "alert-danger",
                            message: data.error.message,
                        }
                    })
                } else {
                    const d = jwt_decode(data.refreshToken)
                    this.handleJWTChange(data.refreshToken, d.userId, d.email, d.role);
                    window.localStorage.setItem("jwt", JSON.stringify(data.refreshToken))
                    this.props.history.push({
                        pathname: "/",
                    })
                }
            })
    };

    handleJWTChange(jwt, userId, email, role) {
        this.props.handleJWTChange(jwt, userId, email, role);
    }

    hasError(key) {
        return this.state.errors.indexOf(key) !== -1;
    }

    render() {
        return (
            <Fragment>
                <h2>Вход</h2>
                <hr/>
                <Alert alertType={this.state.alert.type}
                       alertMessage={this.state.alert.message}/>

                <form className={"pt-3"} onSubmit={this.handleSubmit}>
                    <Input title={"Имя пользователя"} type={"text"} name={"userName"} handleChange={this.handleChange}
                           className={this.hasError("userName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("userName") ? "text-danger" : "d-none"}
                           errorMsg={"Имя пользователя должно быть больше 2, но меньше 33"}/>

                    <Input title={"Пароль"} type={"password"} name={"password"} handleChange={this.handleChange}
                           className={this.hasError("password") ? "is-invalid" : ""}
                           errorDiv={this.hasError("password") ? "text-danger" : "d-none"}
                           errorMsg={"Пароль должен быть больше 2, но меньше 33"}/>

                    <hr/>

                    <button className={"btn btn-primary"}>Войти</button>
                </form>
            </Fragment>
        );
    }
}