import React, { useState, useEffect } from 'react';
import LoadSpinner from '../LoadSpinner';

import IssueForm from "./IssueForm";

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

function IssueModal(props) {
    const [isOpen, setIsOpen] = useState(false);
    const [opacity, setOpacity] = useState(0);

    const [isSubmitting, setIsSubmitting] = useState(false)
    const [error, setError] = useState(null);

    function toggleModal(e) {
        e.stopPropagation();
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

    function onSubmit(issue) {
        setIsSubmitting(true);
        console.log("Heeloo");
        props.onSubmit(issue)
            .then(
                (result) => {
                    setIsSubmitting(false)
                    //close the modal
                    setIsOpen(false);
                    props.onDone();
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

    const submitLabel = props.issue == null ? "Create Issue" : "Update Issue"
    const toggleLabel = props.issue == null ? "New Issue" : props.issue.name

    return (
        <div
        // style={{
        //     "padding": "1.5em"
        // }}
        >
            <button onClick={toggleModal} className="btn btn-primary">{toggleLabel}</button>
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
                        return <IssueForm buttonHandler={onSubmit} buttonLabel={submitLabel} issue={props.issue} />
                    }
                })()}

            </StyledModal>
        </div>
    )
}

export default IssueModal;