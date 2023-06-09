import React, {Component, Fragment} from "react";
import {Link} from "react-router-dom";

export default class Genres extends Component {

    state = {
        genres: [],
        isLoaded: false,
        error: null,
    }

    componentDidMount() {
        fetch("http://localhost/api/v1/genres")
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
                        genres: json,
                        isLoaded: true,
                    },
                    (error) => {
                        this.setState({
                            isLoaded: true,
                            error
                        });
                    });
            });
    }

    render() {
        const {genres, isLoaded, error} = this.state;

        if (error) {
            return <div>Ошибка: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Загрузка...</p>
        } else {
            return (
                <Fragment>
                    <h2>Жанры</h2>

                    <div className={"list-group"}>
                        {genres.map((m, i) => (
                            <Link key={i} to={{
                                pathname: `/genre/${m.genre}`,
                                genreName: m.genre,
                            }} className="list-group-item list-group-item-action">{m.genre}</Link>
                        ))}
                    </div>
                </Fragment>
            )
        }
    }
}