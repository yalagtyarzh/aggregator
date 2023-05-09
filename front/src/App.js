import React, {Component, Fragment} from "react";
import {BrowserRouter as Router, Link, Route, Switch} from "react-router-dom";
import Products from "./components/Products";
import Admin from "./components/Admin";
import Genres from "./components/Genres";
import OneProduct from "./components/OneProduct";
import OneGenre from "./components/OneGenre";
import EditProduct from "./components/EditProduct";
import Login from "./components/Login"
import Register from "./components/Register";
import jwt_decode from "jwt-decode";

export default class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            jwt: "",
            email: "",
            role: "",
            userId: "",
        }
        this.handleJWTChange(this.handleJWTChange.bind(this));
    }

    handleJWTChange = (jwt, userId, email, role) => {
        this.setState({jwt: jwt, userId: userId, email: email, role: role});
    }

    componentDidMount() {
        let t = window.localStorage.getItem("jwt");
        if (t) {
            let decoded = jwt_decode(t)
            if (this.state.jwt === "") {
                this.setState({email: decoded.email, role: decoded.role, userId: decoded.userId, jwt: JSON.parse(t)});
            }
        }
    }

    logout = () => {
        this.setState({jwt: "", email: "", role: "", userId: ""});
        window.localStorage.removeItem("jwt");
    }

    render() {
        let loginLink;
        if (this.state.jwt === "") {
            loginLink = <Link to={"/login"}>Login</Link>
        } else {
            loginLink = <Link to={"/logout"} onClick={this.logout}>Logout</Link>
        }

        return (
            <Router>

                <div className="container">
                    <div className="row">
                        <div className={"col mt-3"}>
                            <h1 className="mt-3">
                                Aggregator
                            </h1>
                        </div>
                        <div className={"col mt-3 text-end"}>
                            <span className="me-2">
                                {loginLink}
                            </span>
                            {this.state.jwt === "" && (
                                <Link to={"/register"}>Sign in</Link>)}
                        </div>


                        <hr className="mb-3"/>

                        <div className="row">
                            <div className="col-md-2">
                                <nav>
                                    <ul className="list-group">
                                        <li className="list-group-item">
                                            <Link to="/">Home</Link>
                                        </li>
                                        <li className="list-group-item">
                                            <Link to="/genres">Genres</Link>
                                        </li>

                                        {this.state.role === "Admin" && (
                                            <Fragment>
                                                <li className="list-group-item">
                                                    <Link to="/admin/product/0">Add Product</Link>
                                                </li>
                                                <li className="list-group-item">
                                                    <Link to="/admin">Manage Products</Link>
                                                </li>
                                            </Fragment>
                                        )}
                                    </ul>
                                </nav>
                            </div>

                            <div className="col-md-10">
                                <Switch>
                                    <Route path="/product/:id" component={(props) => (
                                        <OneProduct {...props} jwt={this.state.jwt} role={this.state.role}/>
                                    )}/>

                                    <Route path="/genre/:id" component={OneGenre}/>
                                    <Route exact path="/genres">
                                        <Genres/>
                                    </Route>

                                    <Route exact path={"/login"} component={(props) => <Login {...props}
                                                                                              handleJWTChange={this.handleJWTChange}/>}/>
                                    <Route exact path={"/register"} component={(props) => <Register {...props}
                                                                                                    handleJWTChange={this.handleJWTChange}/>}/>

                                    <Route path="/admin/product/:id" component={(props) => (
                                        <EditProduct {...props} jwt={this.state.jwt} role={this.state.role}/>
                                    )}/>
                                    <Route path="/admin">
                                        <Admin/>
                                    </Route>
                                    <Route path="/">
                                        <Products/>
                                    </Route>
                                </Switch>
                            </div>
                        </div>
                    </div>
                </div>
            </Router>
        );
    }
}




