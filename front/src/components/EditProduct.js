import {Component, Fragment} from "react";
import "./EditProduct.css";
import MultipleCheckbox from "./form-components/MultipleCheckbox";
import TextArea from "./form-components/TextArea";
import Input from "./form-components/Input";
import Select from "./form-components/Select"
import Alert from "./ui-components/Alert";
import {Link} from "react-router-dom";
import {confirmAlert} from "react-confirm-alert"
import 'react-confirm-alert/src/react-confirm-alert.css'

function isValidUrl(string) {
    try {
        new URL(string);
        return true;
    } catch (err) {
        return false;
    }
}

export default class EditProduct extends Component {
    state = {
        genres: [],
        product: {
            genres: null
        },
        isLoaded: false,
        error: null,
        errors: null,
    };

    constructor(props) {
        super(props);
        this.state = {
            genres: [],
            product: {
                id: 0,
                title: "",
                description: "",
                genres: [],
                studio: "",
                rating: "",
                year: "",
                imageLink: ""
            },
            isLoaded: false,
            error: null,
            errors: [],
            alert: {
                type: "d-none",
                message: "",
            },
            ratingOptions: [
                {id: "0+", value: "0+"},
                {id: "6+", value: "6+"},
                {id: "12+", value: "12+"},
                {id: "16+", value: "16+"},
                {id: "18+", value: "18+"},
            ],
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit = (evt) => {
        evt.preventDefault();

        // client side validation
        let errors = [];
        if (this.state.product.title === "") {
            errors.push("title")
        }

        if (!isValidUrl(this.state.product.imageLink)) {
            errors.push("imageLink")
        }

        this.setState({errors: errors})

        if (errors.length > 0) {
            return false
        }

        const p = this.state.product
        p.year = parseInt(p.year)
        if (p.id === 0) {

            p.id = undefined
        }

        const headers = new Headers()
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + this.props.jwt);

        const requestOptions = {
            method: "POST",
            body: JSON.stringify(p),
            headers: headers,
        }

        let url = 'http://localhost:81/api/v1/admin/product/update'
        if (p.id === undefined) {
            url = 'http://localhost:81/api/v1/admin/product/create'
        }

        fetch(url, requestOptions)
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    const a = {type: "alert-danger", message: data.error.message}
                    this.setState({
                        alert: a,
                    });
                } else {
                    this.props.history.push({
                        pathname: "/admin"
                    })
                }
            })
    };

    handleChange = (evt) => {
        let value = evt.target.value;
        let name = evt.target.name;
        this.setState((prevState) => ({
            product: {
                ...prevState.product,
                [name]: value,
            }
        }))
    }

    confirmDelete = (e) => {
        confirmAlert({
            title: "Удалить продукт?",
            message: "Вы уверены?",
            buttons: [
                {
                    label: "Да",
                    onClick: () => {

                        const p = {id: this.state.product.id, delete: true}
                        const headers = new Headers()
                        headers.append("Content-Type", "application/json");
                        headers.append("Authorization", "Bearer " + this.props.jwt);
                        const requestOptions = {
                            method: "POST",
                            body: JSON.stringify(p),
                            headers: headers
                        }
                        fetch('http://localhost:81/api/v1/admin/product/update', requestOptions)
                            .then(response => response.json())
                            .then(data => {
                                if (data.error) {
                                    const a = {type: "alert-danger", message: data.error.message}
                                    this.setState({
                                        alert: a,
                                    });
                                } else {
                                    this.props.history.push({
                                        pathname: "/admin",
                                    })
                                }
                            })
                    }
                },
                {
                    label: "No",
                    onClick: () => {
                    }
                }
            ]
        })
    }

    hasError(key) {
        return this.state.errors.indexOf(key) !== -1;
    }

    handleCheckboxChange = (evt) => {
        const value = evt.target.value;
        const isChecked = evt.target.checked;
        this.setState(prevState => {
            const {genres} = prevState.product;
            let updatedGenres;
            if (isChecked) {
                updatedGenres = [...(genres || []), {genre: value}]; // <-- добавляем проверку на null
            } else {
                updatedGenres = genres.filter(genre => genre.genre !== value);
            }
            return {
                product: {
                    ...prevState.product,
                    genres: updatedGenres,
                }
            }
        });
    };

    componentDidMount() {
        if (this.props.role !== "Admin") {
            this.props.history.push({
                pathname: "/login",
            });
            return;
        }

        fetch("http://localhost/api/v1/genres")
            .then((response) => {
                if (response.status !== 200) {
                    let err = Error;
                    err.Message = "Invalid response code: " + response.status;
                    this.setState({error: err});
                }
                return response.json()
            })
            .then((json) => {
                this.setState(
                    {
                        genres: json,
                    },
                )
            })

        const id = this.props.match.params.id;
        if (id > 0) {
            fetch("http://localhost/api/v1/products/get/?pid=" + id)
                .then((response) => {
                    if (response.status !== 200) {
                        let err = Error;
                        err.Message = "Invalid response code: " + response.status;
                        this.setState({error: err});
                    }
                    return response.json()
                })
                .then((json) => {
                    this.setState(
                        {
                            product: {
                                id: parseInt(id),
                                title: json.title,
                                description: json.description,
                                year: json.year,
                                studio: json.studio,
                                rating: json.rating,
                                imageLink: json.imageLink,
                            },
                            isLoaded: true,
                        },
                        (error) => {
                            this.setState({
                                isLoaded: true,
                                error,
                            })
                        }
                    )

                    json.genres.forEach((obj) => {
                        document.getElementById(obj.genre).click()
                    });
                })
        } else {
            this.setState({isLoaded: true});
        }
    };

    render() {
        let {product, isLoaded, error} = this.state;

        if (error) {
            return <div>Error: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>Добавить/редактировать продукт</h2>
                    <Alert
                        alertType={this.state.alert.type} alertMessage={this.state.alert.message}/>
                    <hr/>
                    <form onSubmit={this.handleSubmit}>
                        <input type="hidden" name={"id"} id={"id"} value={product.id} onChange={this.handleChange}/>

                        <Input title={"Наименование"} className={this.hasError("title") ? "is-invalid" : ""} type={"text"}
                               name={"title"} value={product.title} errorMsg={"Введите наименование продукта"}
                               errorDiv={this.hasError("title") ? "text-danger" : "d-none"}
                               handleChange={this.handleChange}/>
                        <Input title={"Год выпуска"} type={"number"} name={"year"} value={product.year}
                               handleChange={this.handleChange}/>
                        <Input title={"Студия/Компания"} type={"text"} name={"studio"} value={product.studio}
                               handleChange={this.handleChange}/>

                        <Select title={"Возрастной рейтинг"} name={"rating"} options={this.state.ratingOptions}
                                value={product.rating}
                                handleChange={this.handleChange} placeholder={"Выбор..."}/>

                        <TextArea title={"Описание"} handleChange={this.handleChange} name={"description"}
                                  value={product.description} rows={"3"}/>

                        <Input title={"Ссылка на картинку"} className={this.hasError("imageLink") ? "is-invalid" : ""} type={"text"}
                               name={"imageLink"} value={product.imageLink} errorMsg={"Введите валидную ссылку"}
                               errorDiv={this.hasError("imageLink") ? "text-danger" : "d-none"}
                               handleChange={this.handleChange}/>

                        <MultipleCheckbox
                            title={"Жанры"}
                            name={"genres"}
                            options={this.state.genres}
                            handleChange={this.handleCheckboxChange}
                        />

                        <hr/>

                        <button className={"btn btn-primary"}>Сохранить</button>
                        <Link to={"/admin"} className={"btn btn-warning ms-1"}>
                            Отмена
                        </Link>

                        {product.id > 0 && (
                            <a href={"#!"} onClick={() => this.confirmDelete()}
                               className={"btn btn-danger ms-1"}>
                                Удалить
                            </a>
                        )}
                    </form>
                </Fragment>
            )
        }
    }
}