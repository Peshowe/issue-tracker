import React, { useState, useEffect } from 'react';
import LoadSpinner from '../LoadSpinner';

import Modal from "styled-react-modal";

const StyledModal = Modal.styled`
  width: 30rem;
  height: 20rem;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: grey;
  opacity: ${(props) => props.opacity};
  transition : all 0.3s ease-in-out;`;

function CreateProjectModal(props) {
    const [isOpen, setIsOpen] = useState(false);
    const [opacity, setOpacity] = useState(0);
    const [projectName, setProjectName] = useState("")
    const [isSubmitting, setIsSubmitting] = useState(false)
    const [error, setError] = useState(null);

    function toggleModal(e) {
        setOpacity(0);
        setIsOpen(!isOpen);
    }

    function afterOpen() {
        setTimeout(() => {
            setOpacity(1);
        }, 100);
    }

    function beforeClose() {
        return new Promise((resolve) => {
            setOpacity(0);
            setTimeout(resolve, 300);
        });
    }

    function handleChange(event) {
        setProjectName(event.target.value)
    }

    function postProject() {
        setIsSubmitting(true)
        fetch('/v1/projects', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: projectName
            })
        })
            // .then(res => res.json)
            .then(
                (result) => {
                    setIsSubmitting(false);
                    //close the modal
                    setIsOpen(false);
                    props.onCreate();

                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    setIsSubmitting(false)
                    setError(error);
                }
            );


    }

    return (
        <div
            style={{
                "padding": "1.5em"
            }}
        >
            <button onClick={toggleModal} className="btn btn-primary">New Project</button>
            <StyledModal
                isOpen={isOpen}
                afterOpen={afterOpen}
                beforeClose={beforeClose}
                onBackgroundClick={toggleModal}
                onEscapeKeydown={toggleModal}
                opacity={opacity}
                backgroundProps={{ opacity }}
            >
                {(() => {

                    if (isSubmitting) {
                        return <LoadSpinner />
                    } else if (error) {
                        return <div>Error: {error.message}</div>;
                    } else {
                        return (<div>
                            <form>
                                <label>Project Name: </label>
                                <input type="text" value={projectName} onChange={handleChange} />
                            </form>

                            <button onClick={postProject} className="btn btn-primary">Create</button>
                        </div>)
                    }
                })()}

            </StyledModal>
        </div>
    )
}

export default CreateProjectModal