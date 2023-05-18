import React, {Component} from "react";
import ReviewForm from "./form-components/ReviewForm";
import {confirmAlert} from "react-confirm-alert"
import 'react-confirm-alert/src/react-confirm-alert.css'

export default class Reviews extends Component {
    state = {reviews: [], isLoaded: false, error: null, found: false};

    constructor(props) {
        super(props);

        this.state = {reviews: [], isLoaded: false, error: null, found: false};
    }

    componentDidMount() {
        fetch("http://localhost/api/v1/reviews/get?pid=" + this.props.id)
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
                        reviews: json,
                        isLoaded: true,
                    },
                    (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        })
                    }
                );
                if (json.find(element => element.userId === this.props.userId)) {
                    this.setState({found: true})
                }
            });
    }

    confirmDelete = (id) => {
        confirmAlert({
            title: "Удалить отзыв?",
            message: "Вы уверены?",
            buttons: [
                {
                    label: "Да",
                    onClick: () => {

                        const p = {id: id, delete: true}
                        const headers = new Headers()
                        headers.append("Content-Type", "application/json");
                        headers.append("Authorization", "Bearer " + this.props.jwt);
                        const requestOptions = {
                            method: "POST",
                            body: JSON.stringify(p),
                            headers: headers
                        }
                        fetch("http://localhost/api/v1/reviews/update", requestOptions)
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
                    }
                },
                {
                    label: "No",
                    onClick: () => {
                    }
                }
            ]
        })
    }

    render() {
        return (
            <div>
                {(this.props.jwt !== "" && !this.state.found) && (
                    <ReviewForm pid={this.props.pid} jwt={this.props.jwt}/>
                )}

                {this.state.reviews.length > 0 ? (
                    <div>
                        <h3>Пользовательские отзывы</h3>
                        {this.state.reviews.map((review) => (
                            <div className="card mt-2" key={review.id}>
                                <div className="card-body">
                                    <span className={"d-flex justify-content-between"}>
                                        <h5 className="card-title">{review.userName}</h5>
                                        <p className={"m-0"}>Оценка: {review.score}</p>
                                    </span>
                                    <h6 className="card-subtitle mb-2 text-muted">{review.firstName} {review.lastName}</h6>

                                    <p className="card-text">{review.content}</p>
                                    <div className={"d-flex justify-content-end"}>
                                        {((this.props.userId === review.userId) || (this.props.role !== "Registered")) && (
                                            <a href={"#!"} onClick={() => this.confirmDelete(review.id)}
                                               className={"btn btn-danger ms-1"}>
                                                Удалить
                                            </a>
                                        )}
                                    </div>
                                </div>
                            </div>
                        ))}
                        <hr/>
                    </div>


                ) : (<p>Для данного продукта еще не написан отзыв.</p>)}
            </div>
        );
    };
};
