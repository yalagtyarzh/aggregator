import React, {Component} from "react";
import ReviewForm from "./form-components/ReviewForm";

export default class Reviews extends Component {
    state = {reviews: [], isLoaded: false, error: null, found: false};

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

    render() {
        return (
            <div>
                {(this.props.jwt !== "" && !this.state.found) && (
                    <ReviewForm jwt={this.props.jwt}/>
                )}

                {this.state.reviews.length > 0 ? (
                    <div>
                        <h3>User Reviews</h3>
                        {this.state.reviews.map((review) => (
                            <div className="card mt-2" key={review.id}>
                                <div className="card-body">
                                    <span className={"d-flex justify-content-between"}>
                                        <h5 className="card-title">{review.userName}</h5>
                                        <p className={"m-0"}>Score: {review.score}</p>
                                    </span>
                                    <h6 className="card-subtitle mb-2 text-muted">{review.firstName} {review.lastName}</h6>

                                    <p className="card-text">{review.content}</p>
                                    <div className={"d-flex justify-content-end"}>
                                        {((this.props.userId === review.userId) || (this.props.role !== "Registered")) && (
                                            <a href={"#!"} onClick={() => this.confirmDelete()}
                                               className={"btn btn-danger ms-1"}>
                                                Delete
                                            </a>
                                        )}

                                        {this.props.userId === review.userId && (
                                            <a href={"#!"} onClick={() => this.confirmDelete()}
                                               className={"btn btn-primary ms-1"}>
                                                Edit
                                            </a>
                                        )}
                                    </div>
                                </div>
                            </div>
                        ))}
                        <hr/>
                    </div>


                ) : (<p>No reviews yet.</p>)}
            </div>
        );
    };
};