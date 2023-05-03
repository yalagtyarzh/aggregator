import {Fragment} from "react";

const MultipleCheckbox = (props) => {
    return (
        <div className={"mb-3 d-flex flex-wrap me-2"}>
            {props.options.map((option, index) => {
                return (
                    <Fragment>
                        <div className="form-check mb-2 me-2">
                            <input className="form-check-input" key={index} type="checkbox" value={option.genre}
                                   onChange={props.handleChange} id={option.genre} />
                            <label className="form-check-label" htmlFor="flexCheckDefault">
                                {option.genre}
                            </label>
                        </div>

                    </Fragment>
                )
            })}
        </div>
    );
}

export default MultipleCheckbox