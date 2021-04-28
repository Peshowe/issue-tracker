import React, { useState, useEffect } from 'react';
import LoadSpinner from '../LoadSpinner';

import CreateProjectModal from './CreateProject'

import { Link } from 'react-router-dom';





function ProjectExplorer() {
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [projects, setProjects] = useState([]);

    // Note: the empty deps array [] means 
    // this useEffect will run once
    // similar to componentDidMount()
    function fetchProjects() {
        fetch("/v1/projects")
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    setProjects(result.projects);
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }
    useEffect(() => {
        fetchProjects()
    }, [])

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
        return <LoadSpinner />;
    } else {
        return (
            <div>

                <CreateProjectModal onCreate={fetchProjects} />
                {projects.map(item => (

                    <div className="listItem"><Link to={`/projects/${item.id}`} className="btn btn-primary">{item.name}</Link></div>


                ))}

            </div>
        );
    }
}


export default ProjectExplorer;