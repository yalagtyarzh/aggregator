import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";


export default class OneGenre extends Component {
    state = {
        products: [],
        isLoaded: false,
        error: null,
        genreName: "",
    }

    componentDidMount() {
        fetch("http://localhost/api/v1/products?genre=" + this.props.match.params.id)
            .then((response) => {
                if (response.status !== 200) {
                    let err = Error;
                    err.message = "Invalid response code: " + response.status;
                    this.setState({error: err});
                }
                return response.json()
            })
            .then((json) => {
                this.setState({
                        products: json,
                        isLoaded: true,
                        genreName: this.props.location.genreName,
                    },
                    (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        })
                    }
                );
            });
    }

    render() {
        let {products, isLoaded, error, genreName} = this.state;

        if (!products) {
            products = [];
        }

        if (error) {
            return <div>Ошибка: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Загрузка...</p>
        } else {
            return (
                <Fragment>
                    <h2>Жанр: {genreName}</h2>

                    <div className={"list-group"}>
                        {products.map((m, index) => (
                            <Link key={index} to={`/product/${m.id}`}
                                  className="list-group-item list-group-item-action">
                                <strong>{m.title}</strong><br/>
                                <small className={"text-muted"}>
                                    {m.year}
                                </small>
                                <br/>
                                {m.description.slice(0, 100)}...</Link>
                        ))}
                    </div>
                </Fragment>
            )
        }
    }
}