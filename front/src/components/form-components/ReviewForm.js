import React, { useState } from "react";

const ReviewForm = ({ onSubmit }) => {
    const [comment, setComment] = useState("");
    const [rating, setRating] = useState(0);

    const handleSubmit = (event) => {
        event.preventDefault();
        console.log(comment, rating)
    };

    return (
        <div>
            <h3>Write a review</h3>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="comment">Review:</label>
                    <textarea
                        className="form-control"
                        id="comment"
                        rows="3"
                        placeholder="Enter your comment"
                        value={comment}
                        onChange={(event) => setComment(event.target.value)}
                    />
                </div>
                <div className={"mt-3"}></div>

                <div className="form-group d-flex">
                    <label htmlFor="rating" className={"p-2"}>Rating:</label>
                    <input
                        type="range"
                        className="form-control-range"
                        id="rating"
                        min="0"
                        max="100"
                        value={rating}
                        onChange={(event) => setRating(event.target.value)}
                    />
                    <output className="form-control-range-output p-2">{rating}</output>
                    <button type="submit" className="btn btn-primary ms-auto">
                        Submit
                    </button>
                </div>
            </form>

            <hr/>
        </div>
    );
};

export default ReviewForm;