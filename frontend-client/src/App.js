import './App.css';

import ToggleMenu from './ToggleMenu/ToggleMenu'

import React, { useState, useEffect } from 'react';
import styled from "styled-components";

import { BrowserRouter, Switch, Route } from 'react-router-dom';


import ProjectExplorer from './Project/ProjectExplorer';
import Project from './Project/Project';

import Login from './Login/Login';

import { ModalProvider, BaseModalBackground } from "styled-react-modal";

const FadingBackground = styled(BaseModalBackground)`
  opacity: ${(props) => props.opacity};
  transition: all 0.3s ease-in-out;
`;

const App = () => {

  const [user, setUser] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("auth/user")
      .then(res => res.json())
      .then(
        (result) => {
          setUser(result["user"]);
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
    <div>
      <ModalProvider backgroundComponent={FadingBackground}>
        <ToggleMenu user={user} />
        <BrowserRouter>
          <Switch>
            <Route path="/projects/:projectId" component={Project} />
            <Route path="/login" component={Login} />
            <Route path="/">
              <ProjectExplorer />
            </Route>

          </Switch>
        </BrowserRouter>
      </ModalProvider>
    </div>


  )
}

export default App;
