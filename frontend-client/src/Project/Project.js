import React, { useState, useEffect } from 'react';
import LoadSpinner from '../LoadSpinner';
import IssuesBoard from '../Issue/IssuesBoard'

function Project(props) {
    const [error, setError] = useState(null);
    const [projectLoaded, setProjectLoaded] = useState(false);
    const [project, setProject] = useState(null);

    function fetchProject() {
        fetch(`/v1/projects/byid/${props.match.params.projectId}`)
            .then(res => res.json())
            .then(
                (result) => {
                    setProject(result);
                    setProjectLoaded(true);
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setProjectLoaded(true);
                    setError(error);
                }
            )
    }



    // Note: the empty deps array [] means 
    // this useEffect will run once
    // similar to componentDidMount()
    useEffect(() => {
        fetchProject();
    }, [])

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (!projectLoaded) {
        return <LoadSpinner />;
    } else {
        return (
            <div>
                <h2 style={{
                    "padding-top": "0.5em",
                    "padding-left": "1em"
                }}>{project.name}</h2>
                <IssuesBoard projectId={project.id} user={props.user} />
            </div>
        );
    }
}

export default Project