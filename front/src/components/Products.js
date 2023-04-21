import React, {Component, Fragment} from "react";

export default class Products extends Component {

    state = {products: []};

    componentDidMount() {
        this.setState({
            movies: [
                {id: 1, title: "Mock 1"},
                {id: 1, title: "Mock 2"},
                {id: 1, title: "Mock 3"},
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
                            {m.title}
                        </li>
                    ))}
                </ul>
            </Fragment>
        );
    }
}