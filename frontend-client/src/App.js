// const { ReactDraggable: Draggable, React, ReactDOM } = window;

import logo from './logo.svg';
import './App.css';


import ToggleMenu from './ToggleMenu/ToggleMenu'

import React, { useState, useEffect } from 'react';
import styled from "styled-components";

import { BrowserRouter, Switch, Route } from 'react-router-dom';


import ProjectExplorer from './Project/ProjectExplorer';
import Project from './Project/Project';

import { ModalProvider, BaseModalBackground } from "styled-react-modal";

const FadingBackground = styled(BaseModalBackground)`
  opacity: ${(props) => props.opacity};
  transition: all 0.3s ease-in-out;
`;

const App = () => {

  return (
    <div>
      <ModalProvider backgroundComponent={FadingBackground}>
        <ToggleMenu />
        <BrowserRouter>
          <Switch>
            <Route path="/projects/:projectId" component={Project} />

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
