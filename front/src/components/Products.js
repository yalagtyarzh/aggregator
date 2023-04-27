import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";

export default class Products extends Component {

    state = {products: [], isLoaded: false};

    componentDidMount() {
        fetch("http://user-api/api/v1/products")
            .then((response) => response.json())
            .then((json) => {
                this.setState({
                    products: json,
                    isLoaded: true
                })
            })
    }

    render() {
        return (
            <Fragment>
                <h2>Products</h2>

                <ul>
                    {this.state.products.map((m) => (
                        <li key={m.id}>
                            <Link to={`/product/${m.id}`}>{m.title}</Link>
                        </li>
                    ))}
                </ul>
            </Fragment>
        );
    }
}