import {Component, Fragment} from "react";
import "./EditProduct.css"

export default class EditProduct extends Component {
    state = {
        product: {},
        isLoaded: false,
        error: null,
    }

    render() {
        let {product} = this.state;
        return (
            <Fragment>
                <h2>Add/Edit Product</h2>
                <hr/>
                <form method={"post"}>
                    <div className={"mb-3"}>
                        <label htmlFor="title" className={"form-label"}>
                            Title
                        </label>
                        <input type="text" className={"form-control"} id={"title"} name={"title"} value={product.title}/>
                    </div>
                    <div className={"mb-3"}>
                        <label htmlFor="year" className={"form-label"}>
                            Year
                        </label>
                        <input type="number" className={"form-control"} id={"year"} name={"year"} value={product.year}/>
                    </div>
                    <div className={"mb-3"}>
                        <label htmlFor="studio" className={"form-label"}>
                            Studio
                        </label>
                        <input type="text" className={"form-control"} id={"studio"} name={"studio"} value={product.year}/>
                    </div>
                    <div className={"mb-3"}>
                        <label htmlFor="rating" className={"form-label"}>
                            Rating
                        </label>
                        <input type="text" className={"form-control"} id={"rating"} name={"rating"} value={product.year}/>
                    </div>
                </form>
            </Fragment>
        )
    }
}