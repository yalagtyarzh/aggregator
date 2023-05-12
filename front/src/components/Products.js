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
                    <h2>Choose product</h2>

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
            );
        }
    }
}