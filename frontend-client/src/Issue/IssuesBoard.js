import React, { useState, useEffect } from 'react';
import Draggable from 'react-draggable';
import LoadSpinner from '../LoadSpinner';
import CreateIssueModal from './CreateIssue';

function IssuesBoard(props) {

    const [error, setError] = useState(null);
    const [issuesLoaded, setIssuesLoaded] = useState(false);
    // const [issues, setIssues] = useState([]);

    const [toDo, setToDo] = useState([]);
    const [inProgress, setInProgress] = useState([]);
    const [done, setDone] = useState([]);

    const statusCols = ["todo_col", "inprogress_col", "done_col"];
    const statusVals = ["to do", "in progress", "done"];


    function popIssue(issue) {
        if (issue.status == "to do") {
            setToDo(toDo.filter(is => is.id != issue.id));
        } else if (issue.status == "in progress") {
            setInProgress(inProgress.filter(is => is.id != issue.id));
        } else if (issue.status == "done") {
            setDone(done.filter(is => is.id != issue.id));
        }
    }

    function pushIssue(issue) {

        if (issue.status == "to do") {
            setToDo([...toDo, issue]);
        } else if (issue.status == "in progress") {
            setInProgress([...inProgress, issue]);
        } else if (issue.status == "done") {
            setDone([...done, issue]);
        }
    }


    function putStatus(issue, status) {

        setIssuesLoaded(false)
        fetch('/v1/issues/status', {
            method: 'PUT',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                issueId: issue.id,
                newStatus: status

            })
        })
            // .then(res => res.json)
            .then(
                (result) => {
                    popIssue(issue);
                    issue.status = status;
                    pushIssue(issue);
                    setIssuesLoaded(true);
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setIssuesLoaded(true);
                    setError(error);
                }
            );
    }


    function onStop(e, position, issue) {

        //check in which column we've dragged the issue
        let i;
        for (i = 0; i < statusCols.length; i++) {
            let bounds = document.getElementById(statusCols[i]).getBoundingClientRect();

            if (position.x > bounds.left && position.x < bounds.right) {
                console.log("In " + statusVals[i]);
                console.log(position)
                console.log(bounds)


                if (issue.status == statusVals[i]) {
                    //TODO: reset if not in a known column or in same column
                    // position.x = position.lastX
                    // position.y = position.lastY
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
            draggableIssues.push(<li key={value.id}><Draggable bounds="#board" grid={[50, 50]} onStop={(e, position) => onStop(e, position, value)} ><button>{value.name}</button></Draggable></li>);
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
            <div>

                <CreateIssueModal projectId={props.projectId} onCreate={fetchIssues} />

                    Hello

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