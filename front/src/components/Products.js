import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";

export default class Products extends Component {

    state = {products: []};

    componentDidMount() {
        this.setState({
            products: [
                {id: 1, title: "Mock 1"},
                {id: 2, title: "Mock 2"},
                {id: 3, title: "Mock 3"},
            ]
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