import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";

export default class Products extends Component {

    state = {products: [], isLoaded: false, error: null};

    componentDidMount() {
        fetch("http://localhost/api/v1/products")
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
                        isLoaded: true
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
        const {products, isLoaded, error} = this.state
        if (error) {
            return <div>Error: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>Products</h2>

                    <ul>
                        {products.map((m) => (
                            <li key={m.id}>
                                <Link to={`/product/${m.id}`}>{m.title}</Link>
                            </li>
                        ))}
                    </ul>
                </Fragment>
            );
        }
    }
}