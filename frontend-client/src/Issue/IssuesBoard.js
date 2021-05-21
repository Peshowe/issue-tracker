import React, { useState, useEffect } from 'react';
import Draggable from 'react-draggable';
import LoadSpinner from '../LoadSpinner';
import IssueModal from './IssueModal';
import AddUserModal from '../Project/AddUser';
import LeaveProjectModal from '../Project/LeaveProject';

function IssuesBoard(props) {

    const [error, setError] = useState(null);
    const [issuesLoaded, setIssuesLoaded] = useState(false);

    //arrays where the issues will be stored depending on their statuses
    const [toDo, setToDo] = useState([]);
    const [inProgress, setInProgress] = useState([]);
    const [done, setDone] = useState([]);

    const statusCols = ["todo_col", "inprogress_col", "done_col"];
    const statusVals = ["to do", "in progress", "done"];


    //pop the given issue from its current status array
    function popIssue(issue) {
        if (issue.status == "to do") {
            setToDo(toDo.filter(is => is.id != issue.id));
        } else if (issue.status == "in progress") {
            setInProgress(inProgress.filter(is => is.id != issue.id));
        } else if (issue.status == "done") {
            setDone(done.filter(is => is.id != issue.id));
        }
    }

    //push the given issue into a status array based on its current status
    function pushIssue(issue) {

        if (issue.status == "to do") {
            setToDo([...toDo, issue]);
        } else if (issue.status == "in progress") {
            setInProgress([...inProgress, issue]);
        } else if (issue.status == "done") {
            setDone([...done, issue]);
        }
    }

    function postIssue(issue) {
        return fetch('/v1/issues', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                issue: {
                    name: issue.name,
                    desc: issue.desc,
                    issue_type: issue.issue_type,
                    bug_trace: issue.bug_trace,
                    status: "to do",
                    project: props.projectId,
                    user: issue.user
                }

            })
        })
    }

    function putIssue(issue) {
        return fetch('/v1/issues', {
            method: 'PUT',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                issue: issue

            })
        })
    }

    function putStatus(issue, status) {
        popIssue(issue);
        issue.status = status;
        pushIssue(issue);

        setIssuesLoaded(false);
        putIssue(issue)
            .then(
                (result) => {

                    setIssuesLoaded(true);
                },
                (error) => {
                    setIssuesLoaded(true);
                    setError(error);
                }
            );
    }


    //handler when we stop dragging an issue
    function onStop(e, issue) {
        //check in which column we've dragged the issue
        e.stopPropagation(); //stop other handlers from triggering
        let i;
        for (i = 0; i < statusCols.length; i++) {

            let bounds = document.getElementById(statusCols[i]).getBoundingClientRect();

            if ((e.clientX > bounds.left) && (e.clientX < bounds.right)) {
                console.log("In " + statusVals[i]);

                // putStatus(issue, statusVals[i]);

                if (issue.status == statusVals[i]) {
                    //TODO: reset if not in a known column or in same column
                    setToDo(toDo);
                    console.log("no change");
                } else {
                    // issue.status = statusVals[i];
                    putStatus(issue, statusVals[i]);
                }
                break;
            }
        }
    };



    function fetchIssues() {
        fetch(`/v1/issues/byproject/${props.projectId}`)
            .then(res => res.json())
            .then(
                (result) => {
                    setIssuesLoaded(true);

                    if (result["issues"] != undefined) {
                        let issues = result["issues"];
                        issues.sort((a, b) => a.last_modified_on - b.last_modified_on);
                        // setIssues(issues);

                        setToDo(issues.filter(issue => issue.status == "to do"));
                        setInProgress(issues.filter(issue => issue.status == "in progress"));
                        setDone(issues.filter(issue => issue.status == "done"));
                    }
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setIssuesLoaded(true);
                    setError(error);
                }
            )
    }

    function createDraggableIssues(issues) {
        const draggableIssues = []
        for (const [index, value] of issues.entries()) {
            draggableIssues.push(<Draggable bounds="#board" onDrag={(e) => onStop(e, value)} ><div className="listItem"><IssueModal issue={value} user={props.user} users={props.users} onSubmit={putIssue} onDone={() => ""} /></div></Draggable>);
        }

        return draggableIssues;
    }

    useEffect(() => {
        fetchIssues();
    }, [])

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (!issuesLoaded) {
        return <LoadSpinner />;
    } else {

        const toDoComponents = createDraggableIssues(toDo);
        const inProgressComponents = createDraggableIssues(inProgress);
        const doneComponents = createDraggableIssues(done);

        return (
            <div
                style={
                    { "padding": "1.5em" }
                }>
                <div>
                    <IssueModal projectId={props.projectId} user={props.user} users={props.users} onSubmit={postIssue} issue={null} onDone={fetchIssues} />
                    <AddUserModal projectId={props.projectId} />
                    <LeaveProjectModal projectId={props.projectId} user={props.user} />
                </div>
                <br />
                <div id="board" className="row" style={{ position: 'relative', overflow: 'auto', padding: '0' }}>
                    <div id={statusCols[0]} className="column">
                        To Do
                    {toDoComponents}
                    </div>
                    <div id={statusCols[1]} className="column">
                        In progress
                    {inProgressComponents}
                    </div>
                    <div id={statusCols[2]} className="column">
                        Done
                    {doneComponents}
                    </div>
                </div>


            </div >
        );
    }
}


export default IssuesBoard