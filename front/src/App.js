import React, {Component, Fragment} from "react";
import {BrowserRouter as Router, Link, Route, Switch} from "react-router-dom";
import Products from "./components/Products";
import Home from "./components/Home";
import Admin from "./components/Admin";
import Genres from "./components/Genres";
import OneProduct from "./components/OneProduct";
import OneGenre from "./components/OneGenre";
import EditProduct from "./components/EditProduct";
import Login from "./components/Login"


export default class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            jwt: "",
        }
        this.handleJWTChange(this.handleJWTChange.bind(this));
    }

    handleJWTChange = (jwt) => {
        this.setState({jwt: jwt});
    }

    logout = () => {
        this.setState({jwt: ""});
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
                                            <Link to="/products">Products</Link>
                                        </li>
                                        <li className="list-group-item">
                                            <Link to="/genres">Genres</Link>
                                        </li>

                                        {this.state.jwt !== "" && (
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
                                    <Route path="/product/:id" component={OneProduct}/>
                                    <Route path="/products">
                                        <Products/>
                                    </Route>

                                    <Route path="/genre/:id" component={OneGenre}/>
                                    <Route exact path="/genres">
                                        <Genres/>
                                    </Route>

                                    <Route exact path={"/login"} component={(props) => <Login {...props} handleJWTChange={this.handleJWTChange} />} />
                                    <Route exact path={"/register"} component={(props) => <Register {...props} handleJWTChange={this.handleJWTChange} />} />

                                    <Route path="/admin/product/:id" component={EditProduct}/>
                                    <Route path="/admin">
                                        <Admin/>
                                    </Route>
                                    <Route path="/">
                                        <Home/>
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




