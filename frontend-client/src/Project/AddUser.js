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

function AddUserModal(props) {
    const [isOpen, setIsOpen] = useState(false);
    const [opacity, setOpacity] = useState(0);
    const [newUser, setNewUser] = useState("")
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
        setNewUser(event.target.value)
    }

    function postUser() {
        setIsSubmitting(true)
        fetch('/v1/projects/users', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                projectId: props.projectId,
                userId: newUser
            })
        })
            // .then(res => res.json)
            .then(
                (result) => {
                    setIsSubmitting(false);
                    //close the modal
                    setIsOpen(false);
                    // props.onCreate();

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
        <

            >
            <button onClick={toggleModal} className="btn btn-primary" style={{ "margin": "0.5em" }}>Add User</button>
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
                                <label>User: </label>
                                <input type="text" value={newUser} onChange={handleChange} />
                            </form>

                            <button onClick={postUser} className="btn btn-primary">Add</button>
                        </div>)
                    }
                })()}

            </StyledModal>
        </>
    )
}

export default AddUserModal