import React, {Component, Fragment} from "react";

export default class Users extends Component {
    state = {
        users: [],
        isLoaded: false,
        error: null,
    };

    constructor(props) {
        super(props);
        this.state = {
            users: [],
            isLoaded: false,
            error: null,
        };
    }

    componentDidMount() {
        if (this.props.role !== "Admin") {
            this.props.history.push({
                pathname: "/login",
            });
            return;
        }

        const h = new Headers()
        h.append("Authorization", "Bearer " + this.props.jwt)
        const ro = {
            headers: h,
        }

        fetch("http://localhost:81/api/v1/admin/users", ro)
            .then(resp => resp.json())
            .then(data => {
                if (data.error) {
                    let err = Error;
                    err.message = data.error
                    this.setState({error: err});
                } else {
                    this.setState({users: data, isLoaded: true})
                }
            }, (error) => {
                this.setState({
                    isLoaded: true,
                    error
                })
            })
    }

    promote = (userId) => {

    }

    render() {
        let {users, isLoaded, error} = this.state;

        if (!users) {
            users = [];
        }

        if (error) {
            return <div>Error: {error.message}</div>
        } else if (!isLoaded) {
            return <p>Loading...</p>
        } else {
            return (
                <Fragment>
                    <h2>Users</h2>
                    <hr/>
                    <table className={"table"}>
                        <thead>
                        <tr>
                            <th className={"col-md-5"}>User ID</th>
                            <th>Name</th>
                            <th>Role</th>
                            <th>Action</th>
                        </tr>
                        </thead>
                        <tbody>
                        {users.map((user) => (
                            <tr key={user.userId}>
                                <td>{user.userId}</td>
                                <td>
                                    {user.firstName} {user.lastName}
                                </td>
                                <td>{user.role}</td>
                                <td>
                                    {(user.role === "Registered" &&
                                            user.userId !== this.props.userId && (
                                                <button className="btn btn-primary" onClick={() => {
                                                    const req = {userId: user.userId}
                                                    const h = new Headers()
                                                    h.append("Authorization", "Bearer " + this.props.jwt)
                                                    h.append("Content-Type", "application/json")
                                                    const ro = {
                                                        method: "POST",
                                                        headers: h,
                                                        body: JSON.stringify(req),
                                                    }
                                                    fetch('http://localhost:81/api/v1/admin/promote/Moderator', ro)
                                                        .then(response => response.json())
                                                        .then(data => {
                                                            if (data.error) {
                                                                let err = Error;
                                                                err.message = data.error
                                                                this.setState({error: err});
                                                            } else {
                                                                this.componentDidMount()
                                                            }
                                                        })
                                                }}>Promote</button>
                                            )) ||
                                        (user.role === "Moderator" &&
                                            <button className="btn btn-danger" onClick={() => {
                                                const req = {userId: user.userId}
                                                const h = new Headers()
                                                h.append("Authorization", "Bearer " + this.props.jwt)
                                                h.append("Content-Type", "application/json")
                                                const ro = {
                                                    method: "POST",
                                                    headers: h,
                                                    body: JSON.stringify(req),
                                                }
                                                fetch('http://localhost:81/api/v1/admin/promote/Registered', ro)
                                                    .then(response => response.json())
                                                    .then(data => {
                                                        if (data.error) {
                                                            let err = Error;
                                                            err.message = data.error
                                                            this.setState({error: err});
                                                        } else {
                                                            this.componentDidMount()
                                                        }
                                                    })
                                            }}>Dismiss</button>
                                        )}
                                </td>
                            </tr>
                        ))}
                        </tbody>
                    </table>
                </Fragment>
            );
        }
    }
}