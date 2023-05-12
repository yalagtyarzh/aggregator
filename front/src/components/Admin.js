import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";

export default class Admin extends Component {
    state = {
        movies: [],
        isLoaded: false,
        error: null,
    }

    componentDidMount() {
        if (this.props.role !== "Admin") {
            this.props.history.push({
                pathname: "/login",
            });
            return;
        }

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
                    <h2>Manage products</h2>

                    <div className={"list-group"}>
                        {products.map((m, index) => (
                            <Link key={index} to={`/admin/product/${m.id}`}
                                  className="list-group-item list-group-item-action">{m.title}</Link>
                        ))}
                    </div>
                </Fragment>
            );
        }
    }
}