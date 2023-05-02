const TextArea = (props) => {
    return (
        <div className={"mb-3"}>
            <label htmlFor="description" className={"form-label"}>
                {props.title}
            </label>
            <textarea className={"form-control"} onChange={props.handleChange} value={props.value} name={props.name}
                      id={props.name} rows={props.rows}/>
        </div>
    )
}

export default TextArea;