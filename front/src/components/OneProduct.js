import React, {Component, Fragment} from "react";
import "./OneProduct.css"
import Reviews from "./Reviews";

export default class OneProduct extends Component {

    state = {product: {}, isLoaded: false, error: null};

    componentDidMount() {
        fetch("http://localhost/api/v1/products/get/?pid=" + this.props.match.params.id)
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
                        product: json,
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
        const {product, isLoaded, error} = this.state;
        if (product.genres) {
            product.genres = Object.values(product.genres)
        } else {
            product.genres = []
        }

        if (error) {
            return <div>Error: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>
                        {product.title.toUpperCase()} ({product.year})
                    </h2>

                    <div className={"float-start"}>
                        <small>Rating: {product.rating}</small>
                    </div>
                    <div className={"float-end"}>
                        {product.genres.map((m, index) => (
                            <span className={"badge bg-secondary me-1"} key={index}>
                                    {m.genre}
                                </span>
                        ))}
                    </div>
                    <div className={"clearfix"}></div>

                    <hr/>

                    <div className="row d-flex align-items-center">
                        <div className="col-md-3 d-flex justify-content-center align-items-center product-image"
                             style={{height: "200px"}}>
                            <img
                                src={this.state.product.imageLink} className={"rounded border border-secondary"}
                                alt="xd" style={{maxWidth: '100%', height: '100%', objectFit: 'cover'}}/>
                        </div>
                        <div className="col-md-9">
                            <table className="table table-compact">
                                <thead></thead>
                                <tbody>
                                <tr>
                                    <td><strong>Description:</strong></td>
                                    <td>{this.state.product.description}</td>
                                </tr>
                                <tr>
                                    <td><strong>Studio:</strong></td>
                                    <td>{this.state.product.studio}</td>
                                </tr>
                                <tr>
                                    <td><strong>Score:</strong></td>
                                    <td>{this.state.product.score}</td>
                                </tr>
                                </tbody>
                            </table>
                        </div>
                    </div>

                    <hr/>

                    <Reviews  jwt={this.props.jwt} role={this.props.role} id={this.props.match.params.id} userId={this.props.userId} pid={this.props.match.params.id}/>

                </Fragment>
            )
        }
    }
}