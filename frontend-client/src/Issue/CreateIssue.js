import React, { useState, useEffect } from 'react';
import LoadSpinner from '../LoadSpinner';

import Modal from "styled-react-modal";

const StyledModal = Modal.styled`
  width: 70rem;
  height: 40rem;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: grey;
  opacity: ${(props) => props.opacity};
  transition : all 0.3s ease-in-out;`;

function CreateIssueModal(props) {
    const [isOpen, setIsOpen] = useState(false);
    const [opacity, setOpacity] = useState(0);

    const [name, setName] = useState("")
    const [description, setDescription] = useState("")
    const [issueType, setIssueType] = useState("feature")
    const [bugTrace, setBugTrace] = useState("")

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

    function postIssue() {
        setIsSubmitting(true)
        fetch('/v1/issues', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                issue: {
                    name: name,
                    desc: description,
                    issue_type: issueType,
                    bug_trace: bugTrace,
                    status: "to do",
                    project: props.projectId
                }

            })
        })
            // .then(res => res.json)
            .then(
                (result) => {
                    setIsSubmitting(false)
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
            <button onClick={toggleModal} className="btn btn-primary">New Issue</button>
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

                            <button onClick={postIssue} className="btn btn-primary">Create</button>
                        </div>)
                    }
                })()}

            </StyledModal>
        </div>
    )
}

export default CreateIssueModal