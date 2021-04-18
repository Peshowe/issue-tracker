import React, { useState, useEffect } from 'react';


function IssueForm(props) {
    const [issue, setIssue] = useState({});
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [issueType, setIssueType] = useState("feature");
    const [bugTrace, setBugTrace] = useState("");


    function onClick() {
        //wrap the issue's attributes in a single object and pass it to the handler
        issue.name = name;
        issue.desc = description;
        issue.issue_type = issueType;
        issue.bug_trace = bugTrace;
        console.log(issue);
        console.log(description);
        setIssue(issue);
        props.buttonHandler(issue);
    }

    useEffect(() => {
        // if an issue was passed as a prop, use its attributes to populate the fields
        if (props.issue != null) {
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
        }
    }, [])


    return (<div>
        <form>
            <label>Name: </label>
            <input type="text" value={name} onChange={e => setName(e.target.value)} />

            <label>Description: </label>
            <input type="text" value={description} onChange={e => setDescription(e.target.value)} />

            <label>Issue type: </label>
            <select value={issueType} onChange={e => setIssueType(e.target.value)}>
                <option value="bug">Bug</option>
                <option value="feature">Feature</option>
                <option value="adhoc">Ad Hoc</option>
            </select>

            <label>Bug trace: </label>
            <textarea value={bugTrace} onChange={e => setBugTrace(e.target.value)} />
        </form>

        <button onClick={onClick} className="btn btn-primary">{props.buttonLabel}</button>
    </div>)
}

export default IssueForm;