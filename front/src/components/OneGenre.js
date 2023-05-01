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
            return <div>Error: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>Genre: {genreName}</h2>

                    <div className={"list-group"}>
                        {products.map((m) => (

                            <Link to={`/product/${m.id}`}
                                  className="list-group-item list-group-item-action">{m.title}</Link>
                        ))}
                    </div>
                </Fragment>
            )
        }
    }
}