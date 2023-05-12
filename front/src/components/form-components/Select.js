const Select = (props) => {
    return (
        <div className={"mb-3"}>
            <label htmlFor={props.name} className={"form-label"}>
                {" "}
                {props.title}{" "}
            </label>
            <select className={"form-select"} name={props.name} id={props.name} value={props.value}
                    onChange={props.handleChange}>
                <option value="">{props.placeholder}</option>
                {props.options.map((option, index) => {
                    return (
                        <option className={"form-select"} key={index} value={option.value}>{option.value}</option>
                    )
                })}
            </select>
        </div>
    );
}

export default Select