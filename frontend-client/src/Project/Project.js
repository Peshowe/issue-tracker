import React, { useState, useEffect } from 'react';
import LoadSpinner from '../LoadSpinner';

import CreateIssueModal from '../Issue/CreateIssue';

function Project(props) {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [project, setProject] = useState([]);

    // Note: the empty deps array [] means 
    // this useEffect will run once
    // similar to componentDidMount()
    useEffect(() => {
        fetch(`/v1/projects/byid/${props.match.params.projectId}`)
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    setProject(result);
                    console.log(result);
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }, [])

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
        return <LoadSpinner />;
    } else {
        return (
            <div>
                {project.name}
                <CreateIssueModal projectId={project.id} />
            </div>
            // <ul>
            //     {project.map(item => (
            //         <li key={item.id}>
            //             <button variant="contained">{item.name} {item.created_on}</button>
            //         </li>
            //     ))}
            //     <LoadSpinner/>
            // </ul>
        );
    }
}

export default Project