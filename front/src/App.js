import React from "react";
import {BrowserRouter as Router, Link, Route, Switch, useRouteMatch} from "react-router-dom";
import Products from "./components/Products";
import Home from "./components/Home";
import Admin from "./components/Admin";
import Genres from "./components/Genres";
import OneProduct from "./components/OneProduct";


export default function App() {
    return (
        <Router>

            <div className="container">
                <div className="row">
                    <h1 className="mt-3">
                        Aggregator
                    </h1>
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
                                        <Link to="/by-genre">Genres</Link>
                                    </li>
                                    <li className="list-group-item">
                                        <Link to="/admin">Manage Products</Link>
                                    </li>
                                </ul>
                            </nav>
                        </div>

                        <div className="col-md-10">
                            <Switch>
                                <Route path="/product/:id" component={OneProduct}/>
                                <Route path="/products">
                                    <Products/>
                                </Route>
                                <Route exact path="/by-genre">
                                    <GenrePage/>
                                </Route>

                                <Route
                                    exact
                                    path="/by-genre/drama"
                                    render={(props) => <Genres {...props} title={`Drama`}/>}
                                />

                                <Route
                                    exact
                                    path="/by-genre/comedy"
                                    render={(props) => <Genres {...props} title={`Comedy`}/>}
                                />

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

function GenrePage() {

    let {path, url} = useRouteMatch();
    return (
        <div>
            <h2>Genres</h2>

            <ul>
                <li><Link to={`${path}/drama`}>Drama</Link></li>
                <li><Link to={`${url}/comedy`}>Comedy</Link></li>
            </ul>
        </div>
    );
}



