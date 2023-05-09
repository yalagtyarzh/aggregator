import React, {Component} from "react";
import Alert from "../ui-components/Alert";
import {confirmAlert} from "react-confirm-alert";

export default class ReviewForm extends Component {
    state = {
        errors: [],
    };

    constructor(props) {
        super(props);
        this.state = {
            errors: [],
            review: {
                content: "",
                score: 0,
            },
            alert: {
                type: "d-none",
                message: "",
            },
        };

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleSubmit = (evt) => {
        evt.preventDefault();

        let errors = [];

        if (this.state.review.content.length < 20) {
            errors.push("content")
        }

        if (this.state.review.score < 0 || this.state.review.score > 100) {
            errors.push("score")
        }

        this.setState({errors: errors})

        if (errors.length > 0) {
            return false
        }

        const req = {
            productId: Number(this.props.pid),
            score: Number(this.state.review.score),
            content: this.state.review.content,
            contentHTML: this.state.review.content,
        }

        const headers = new Headers()
        headers.append("Content-Type", "application/json");
        headers.append("Authorization", "Bearer " + this.props.jwt);

        const requestOptions = {
            method: "POST",
            body: JSON.stringify(req),
            headers: headers,
        }


        fetch("http://localhost/api/v1/reviews/create", requestOptions)
            .then(response => response.json())
            .then(data => {
                if (data.error) {
                    const a = {type: "alert-danger", message: data.error.message}
                    this.setState({
                        alert: a,
                    });
                } else {
                    window.location.reload()
                }
            })
    };

    handleChange = (evt) => {
        let value = evt.target.value;
        let name = evt.target.name;
        this.setState((prevState) => ({
            review: {
                ...prevState.review,
                [name]: value,
            }
        }))
    }

    hasError(key) {
        return this.state.errors.indexOf(key) !== -1;
    }

    render() {
        let {review} = this.state;

        return (
            <div>
                <h3>Write a review</h3>
                <Alert
                    alertType={this.state.alert.type} alertMessage={this.state.alert.message}/>
                <form onSubmit={this.handleSubmit}>
                    <div className="form-group">
                        <label htmlFor="content">Review:</label>
                        <textarea
                            name={"content"}
                            className={`form-control ${this.hasError("content") ? "is-invalid" : ""}`}
                            id="content"
                            rows="3"
                            placeholder="Enter your comment"
                            onChange={this.handleChange}
                        />
                    </div>
                    <div className={this.hasError("content") ? "text-danger" : "d-none"}>Need at least 20 characters</div>
                    <div className={"mt-3"} />

                    <div className="form-group d-flex">
                        <label htmlFor="score" className={"p-2"}>Score:</label>
                        <input
                            name={"score"}
                            type="range"
                            className={`form-control-range`}
                            id="score"
                            inputMode={"numeric"}
                            defaultValue={0}
                            min="0"
                            max="100"
                            onChange={this.handleChange}
                        />

                        <output className="form-control-range-output p-2">{review.score}</output>
                        <div className={this.hasError("score") ? "text-danger" : "d-none"}>Invalid score</div>
                        <button type="submit" className="btn btn-primary ms-auto">
                            Submit
                        </button>
                    </div>
                </form>

                <hr/>
            </div>
        );
    };
};