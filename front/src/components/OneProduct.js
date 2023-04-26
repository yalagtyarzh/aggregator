import {Component, Fragment} from "react";

export default class OneProduct extends Component {

    state = {product: {}};

    componentDidMount() {
        this.setState({
            product: {
                id: this.props.match.params.id,
                title: "Some product",
                runtime: 150,
            }
        })
    }

    render() {
        return(
            <Fragment>
                <h2>Product: {this.state.product.title}</h2>

                <table className="table table-compact table-striped">
                    <thead></thead>
                    <tbody>
                    <tr>
                        <td><strong>Title:</strong></td>
                        <td>{this.state.product.title}</td>
                    </tr>
                    <tr>
                        <td><strong>Runtime:</strong></td>
                        <td>{this.state.product.runtime} minutes</td>
                    </tr>
                    </tbody>
                </table>
            </Fragment>
        )
    }
}