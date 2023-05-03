import {Component, Fragment} from "react";
import "./EditProduct.css";
import MultipleCheckbox from "./form-components/MultipleCheckbox";
import TextArea from "./form-components/TextArea";
import Input from "./form-components/Input";
import Select from "./form-components/Select"

export default class EditProduct extends Component {
    state = {
        genres: [],
        product: {
            genres: null // <-- заменяем пустой массив на null
        },
        isLoaded: false,
        error: null,
    };

    constructor(props) {
        super(props);
        this.state = {
            genres: [],
            product: {
                id: 0
            },
            isLoaded: false,
            error: null,
            ratingOptions: [
                {id: "G", value: "G"},
                {id: "PG", value: "PG"},
                {id: "PG-13", value: "PG-13"},
                {id: "R", value: "R"},
                {id: "NC-17", value: "NC-17"},
            ],
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit = (evt) => {
        evt.preventDefault();

        const p = this.state.product
        if (p.id === 0) {
            p.id = undefined
        }
        const requestOptions = {
            method: "POST",
            body: JSON.stringify(JSON.stringify(p))
        }

        let url = 'http://localhost:81/api/v1/admin/product/update'
        if (p.id === undefined) {
            url = 'http://localhost:81/api/v1/admin/product/create'
        }

        fetch(url, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
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
                console.log(json)
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
                    console.log(json);
                    this.setState(
                        {
                            product: {
                                id: parseInt(id),
                                title: json.title,
                                description: json.description,
                                year: json.year,
                                studio: json.studio,
                                rating: json.rating,
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
                    <h2>Add/Edit Product</h2>
                    <hr/>
                    <form onSubmit={this.handleSubmit}>
                        <input type="hidden" name={"id"} id={"id"} value={product.id} onChange={this.handleChange}/>

                        <Input title={"Title"} type={"text"} name={"title"} value={product.title}
                               handleChange={this.handleChange}/>
                        <Input title={"Year"} type={"text"} name={"year"} value={product.year}
                               handleChange={this.handleChange}/>
                        <Input title={"Studio"} type={"text"} name={"studio"} value={product.studio}
                               handleChange={this.handleChange}/>

                        <Select title={"Rating"} name={"rating"} options={this.state.ratingOptions}
                                value={product.rating}
                                handleChange={this.handleChange} placeholder={"Choose..."}/>

                        <TextArea title={"Description"} handleChange={this.handleChange} name={"description"}
                                  value={product.description} rows={"3"}/>

                        <MultipleCheckbox
                            title={"Genres"}
                            name={"genres"}
                            options={this.state.genres}
                            handleChange={this.handleCheckboxChange}
                        />

                        <hr/>

                        <button className={"btn btn-primary"}>Save</button>
                    </form>

                    <div className={"mt-3"}>
                        <pre>{JSON.stringify(this.state, null, 3)}</pre>
                    </div>
                </Fragment>
            )
        }
    }
}