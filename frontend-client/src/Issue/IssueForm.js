import React, { useState, useEffect } from 'react';
import Select from 'react-select';


function IssueForm(props) {
    const [issue, setIssue] = useState({});
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [issueType, setIssueType] = useState("feature");
    const [bugTrace, setBugTrace] = useState("");
    const [owner, setOwner] = useState(null);

    function onClick() {
        //wrap the issue's attributes in a single object and pass it to the handler
        issue.name = name;
        issue.desc = description;
        issue.issue_type = issueType;
        issue.bug_trace = bugTrace;
        issue.user = owner;
        setIssue(issue);
        props.buttonHandler(issue);
    }

    useEffect(() => {
        if (props.issue != null) {
            // if an issue was passed as a prop, use its attributes to populate the fields
            setIssue(props.issue)
            if ("name" in props.issue) {
                setName(props.issue.name);
            }
            if ("desc" in props.issue) {
                setDescription(props.issue.desc);
            }
            if ("issue_type" in props.issue) {
                setIssueType(props.issue.issue_type);
            }
            if ("bug_trace" in props.issue) {
                setBugTrace(props.issue.bug_trace);
            }
            if ("user" in props.issue) {
                setOwner(props.issue.user);
            } else {
                setOwner("");
            }
        } else {
            //if no issue was passed, then we're creating a new issue
            // set some default values
            setOwner(props.user)
        }
    }, [])

    const ownerOptions = [];
    props.users.forEach(el => ownerOptions.push({ value: el, label: el }));

    let ownerSelect;

    if (owner != null) {

        ownerSelect = (<Select
            defaultValue={{ "value": owner, "label": owner }}
            options={ownerOptions}
            className="basic-multi-select"
            classNamePrefix="select"
            onChange={e => setOwner(e.value)}
        />)
    }

    return (<div
        style={{
            // "height": "70%",
            // "width": "70%"
        }}
    >
        <form>
            <div>
                <label>Name: </label>
                <input type="text" value={name} onChange={e => setName(e.target.value)} />

                <label>Issue type: </label>
                <select value={issueType} onChange={e => setIssueType(e.target.value)}>
                    <option value="bug">Bug</option>
                    <option value="feature">Feature</option>
                    <option value="adhoc">Ad Hoc</option>
                </select>
            </div>

            <div>
                <label>Description: </label>
                <br />
                <textarea className="inputArea" value={description} onChange={e => setDescription(e.target.value)} />


            </div>
            <div>
                <label>Bug trace: </label>
                <br />
                <textarea className="inputArea" value={bugTrace} onChange={e => setBugTrace(e.target.value)} />
            </div>
            <div>
                <label>Owner: </label>
                {ownerSelect}
                <br />
            </div>
        </form>

        <button onClick={onClick} className="btn btn-primary">{props.buttonLabel}</button>
    </div>)
}

export default IssueForm;