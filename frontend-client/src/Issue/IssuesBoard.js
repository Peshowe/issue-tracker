import React, { useState, useEffect } from 'react';
import Draggable from 'react-draggable';
import LoadSpinner from '../LoadSpinner';
import CreateIssueModal from './CreateIssue';

function IssuesBoard(props) {

    const [error, setError] = useState(null);
    const [issuesLoaded, setIssuesLoaded] = useState(false);
    const [issues, setIssues] = useState([]);

    const statusCols = ["todo_col", "inprogress_col", "done_col"];
    const statusVals = ["to do", "in progress", "done"];


    function onStop(e, position, issue) {

        //check in which column we've dragged the issue
        let i;
        for (i = 0; i < statusCols.length; i++) {
            let bounds = document.getElementById(statusCols[i]).getBoundingClientRect();

            if (position.x > bounds.left && position.x < bounds.right) {
                console.log("In " + statusVals[i]);

                if (issue.status == statusVals[i]) {
                    //TODO: reset if not in a known column or in same column
                    // position.x = position.lastX
                    // position.y = position.lastY
                    console.log("no change");
                } else {
                    issue.status = statusVals[i];

                }
                break;
            }

        }
        console.log(issue);
    };



    function fetchIssues() {
        fetch(`/v1/issues/byproject/${props.projectId}`)
            .then(res => res.json())
            .then(
                (result) => {
                    setIssuesLoaded(true);

                    if (result["issues"] != undefined) {
                        setIssues(result["issues"]);
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
            draggableIssues.push(<Draggable bounds="#board" grid={[50, 50]} onStop={(e, position) => onStop(e, position, value)} ><li key={index}>{value.name}</li></Draggable>);
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

        const toDo = createDraggableIssues(issues.filter(issue => issue.status == "to do"));
        const inProgress = createDraggableIssues(issues.filter(issue => issue.status == "in progress"));
        const done = createDraggableIssues(issues.filter(issue => issue.status == "done"));

        return (
            <div>

                <CreateIssueModal projectId={props.projectId} onCreate={fetchIssues} />

                    Hello

                <div id="board" class="row" style={{ position: 'relative', overflow: 'auto', padding: '0' }}>
                    <div id={statusCols[0]} class="column">
                        To Do
                    {toDo}
                    </div>
                    <div id={statusCols[1]} class="column">
                        In progress
                    {inProgress}
                    </div>
                    <div id={statusCols[2]} class="column">
                        Done
                    {done}
                    </div>
                </div>


            </div >
        );
    }
}


export default IssuesBoard