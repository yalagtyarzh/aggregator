import {Component, Fragment} from "react";
import Alert from "./ui-components/Alert";
import Input from "./form-components/Input";
import jwt_decode from "jwt-decode";

const validateEmail = (email) => {
    return String(email)
        .toLowerCase()
        .match(
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
        );
};

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

        let errors = [];

        if (this.state.firstName === "" || this.state.firstName.length < 2 || this.state.firstName.length > 32) {
            errors.push("firstName")
        }

        if (this.state.lastName === "" || this.state.lastName.length < 2 || this.state.lastName.length > 32) {
            errors.push("lastName")
        }

        if (this.state.userName === "" || this.state.userName.length < 3 || this.state.userName.length > 32) {
            errors.push("userName");
        }

        if (!validateEmail(this.state.email) || this.state.userName.email >150) {
            errors.push("email");
        }

        if (this.state.password === "" || this.state.password.length < 3 || this.state.password.length > 32) {
            errors.push("password");
        }

        if (this.state.password !== this.state.confirm) {
            errors.push("confirm");
        }

        this.setState({errors: errors});

        if (errors.length > 0) {
            return false;
        }

        const req = {
            firstName: this.state.firstName,
            lastName: this.state.lastName,
            userName: this.state.userName,
            email: this.state.email,
            password: this.state.password,
        }

        const requestOptions = {
            method: "POST",
            body: JSON.stringify(req)
        }


        fetch("http://localhost/api/v1/registration", requestOptions)
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
    }

    handleJWTChange(jwt, userId, email, role) {
        this.props.handleJWTChange(jwt, userId, email, role);
    }

    hasError(key) {
        return this.state.errors.indexOf(key) !== -1;
    }

    render() {
        return (
            <Fragment>
                <h2>Регистрация</h2>
                <hr/>
                <Alert alertType={this.state.alert.type}
                       alertMessage={this.state.alert.message}/>


                <form className={"pt-3"} onSubmit={this.handleSubmit}>
                    <Input title={"Имя"} type={"text"} name={"firstName"} handleChange={this.handleChange}
                           className={this.hasError("firstName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("firstName") ? "text-danger" : "d-none"}
                           errorMsg={"Имя должно быть больше 1, но меньше 33 символов"}/>

                    <Input title={"Фамилия"} type={"text"} name={"lastName"} handleChange={this.handleChange}
                           className={this.hasError("lastName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("lastName") ? "text-danger" : "d-none"}
                           errorMsg={"Фамилия должна быть больше 1, но меньше 33 символов"}/>

                    <Input title={"Имя пользователя"} type={"text"} name={"userName"} handleChange={this.handleChange}
                           className={this.hasError("userName") ? "is-invalid" : ""}
                           errorDiv={this.hasError("userName") ? "text-danger" : "d-none"}
                           errorMsg={"Имя пользователя должно быть больше 2, но меньше 33"}/>

                    <Input title={"Электронная почта"} type={"email"} name={"email"} handleChange={this.handleChange}
                           className={this.hasError("email") ? "is-invalid" : ""}
                           errorDiv={this.hasError("email") ? "text-danger" : "d-none"}
                           errorMsg={"Введите валидную электронную почту"}/>

                    <Input title={"Пароль"} type={"password"} name={"password"} handleChange={this.handleChange}
                           className={this.hasError("password") ? "is-invalid" : ""}
                           errorDiv={this.hasError("password") ? "text-danger" : "d-none"}
                           errorMsg={"Пароль должен быть больше 2, но меньше 33"}/>

                    <Input title={"Повторите пароль"} type={"password"} name={"confirm"} handleChange={this.handleChange}
                           className={this.hasError("confirm") ? "is-invalid" : ""}
                           errorDiv={this.hasError("confirm") ? "text-danger" : "d-none"}
                           errorMsg={"Пароли не совпадают"}/>

                    <hr/>

                    <button className={"btn btn-primary"}>Зарегистрироваться</button>
                </form>
            </Fragment>
        );
    }
}