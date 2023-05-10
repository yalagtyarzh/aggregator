import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";
import Input from "./form-components/Input";

export default class GraphQL extends Component {
    constructor(props) {
        super(props);
        this.state = {
            products: [],
            isLoaded: false,
            error: null,
            alert: {
                type: "d-none",
                message: "",
            },
            searchTerm: "",
        }

        this.handleChange = this.handleChange.bind(this);
    }

    handleChange = (evt) => {
        let value = evt.target.value;

        this.setState(
            (prevState) => ({
                searchTerm: value,
            })
        );

        if (value.length > 2) {
            this.performSearch();
        } else {
            this.setState({products: []})
        }

    }

    performSearch() {
        const req = `
        {
            search(titleContains: "${this.state.searchTerm}") {
                id
                title
                year
                description
            }
        }`

        const h = new Headers();
        h.append("Content-Type", "application/json")

        const ro = {
            method: "POST",
            body: req,
            headers: h
        }

        fetch("http://localhost/api/v1/graphql/list", ro)
            .then((response) => response.json())
            .then((data) => {
                return Object.values(data.data.search);
            })
            .then((l) => {
                if (l.length > 0) {
                    this.setState({
                        products: l,
                    })
                } else {
                    this.setState({
                        products: [],
                    })
                }
            })
    }

    componentDidMount() {
        const req = `
        {
            list {
                id
                title
                year
                description
            }
        }`

        const h = new Headers();
        h.append("Content-Type", "application/json")

        const ro = {
            method: "POST",
            body: req,
            headers: h
        }

        fetch("http://localhost/api/v1/graphql/list", ro)
            .then((response) => response.json())
            .then((data) => {
                return Object.values(data.data.list);
            })
            .then((l) => {
                this.setState({
                    products: l,
                })
            })
    }

    render() {
        let {products} = this.state;

        return (
            <Fragment>
                <h2>Search product</h2>
                <hr/>

                <Input
                    title={"Search"} type={"text"} name={"search"} value={this.state.searchTerm} handleChange={this.handleChange}
                />

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