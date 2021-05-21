import React, { useState, useEffect } from 'react';
import styled from 'styled-components'

const Ul = styled.ul`
  list-style-type: none;
  margin: 0;
  padding: 0;
`

const Li = styled.li`
  display: block;
  margin: 0;
  padding: 0;
`

const MyDiv = styled.div`
  border-bottom: 1px solid #efefef;
  color: ${props => props.active ? 'rgba(218, 135, 196)' : '#333'};
  display: block;
  font-family: Helvetica, Arial, sans-serif;
  font-size: 16px;
  font-weight: 700;
  margin-left: 18px;
  padding: 18px 18px 18px 5px;
  text-decoration: none;
  text-transform: uppercase;
  transition: background 200ms ease-in-out;
  &:hover {
    color: rgba(218, 135, 196);
  }
`

function List(props) {

  const [userPreference, setUserPreference] = useState("");
  const [error, setError] = useState(null);

  function putUserPreference(event) {
    setUserPreference(event.target.value)
    fetch('/v1/mailer/preferences', {
      method: 'PUT',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        user: props.user,
        is_mail_notification_on: (event.target.value == "true")

      })
    }).then(
      (error) => {
        setError(error);
      }
    );
  }

  useEffect(() => {
    fetch(`v1/mailer/preferences/${props.user}`)
      .then(res => res.json())
      .then(
        (result) => {
          if (result["is_mail_notification_on"] == null) {
            setUserPreference("false");
          } else {
            setUserPreference(result["is_mail_notification_on"]);
          }
        },
        // Note: it's important to handle errors here
        // instead of a catch() block so that we don't swallow
        // exceptions from actual bugs in components.
        (error) => {
          setError(error);
        }
      )
  }, [])

  return (
    <>
      <Ul>
        <Li>
          <MyDiv>
            Mail notifications:
            <select value={userPreference} onChange={putUserPreference}>
              <option value="true">On</option>
              <option value="false">Off</option>
            </select>
          </MyDiv>
        </Li>
      </Ul>
    </>
  )
}
export default List