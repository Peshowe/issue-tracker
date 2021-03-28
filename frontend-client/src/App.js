// const { ReactDraggable: Draggable, React, ReactDOM } = window;

import logo from './logo.svg';
import './App.css';

import Draggable from 'react-draggable';
import React from 'react';

import Project from './Project/Project'

class App extends React.Component {

  render() {
    return (
      <Project></Project>
    )
  }

  // state = {
  //   activeDrags: 0,
  //   deltaPosition: {
  //     x: 0, y: 0
  //   },
  //   controlledPosition: {
  //     x: -400, y: 200
  //   }
  // };

  // handleDrag = (e, ui) => {
  //   const { x, y } = this.state.deltaPosition;
  //   this.setState({
  //     deltaPosition: {
  //       x: x + ui.deltaX,
  //       y: y + ui.deltaY,
  //     }
  //   });
  // };

  // onStart = () => {
  //   this.setState({ activeDrags: ++this.state.activeDrags });
  // };

  // onStop = () => {
  //   this.setState({ activeDrags: --this.state.activeDrags });
  // };
  // onDrop = (e) => {
  //   this.setState({ activeDrags: --this.state.activeDrags });
  //   if (e.target.classList.contains("drop-target")) {
  //     alert("Dropped!");
  //     e.target.classList.remove('hovered');
  //   }
  // };
  // onDropAreaMouseEnter = (e) => {
  //   if (this.state.activeDrags) {
  //     e.target.classList.add('hovered');
  //   }
  // }
  // onDropAreaMouseLeave = (e) => {
  //   e.target.classList.remove('hovered');
  // }

  // render() {
  //   const dragHandlers = { onStart: this.onStart, onStop: this.onStop };
  //   return (
  //     <div className="App">
  //       <header className="App-header">

  //         <img src={logo} className="App-logo" alt="logo" />

  //         <Draggable {...dragHandlers}>
  //           <div className="box">I can be dragged anywhere</div>
  //         </Draggable>

  //         <p>
  //           Edit <code>src/App.js</code> and save to reload.
  //       </p>
  //         <a
  //           className="App-link"
  //           href="https://reactjs.org"
  //           target="_blank"
  //           rel="noopener noreferrer"
  //         >
  //           Learn React
  //       </a>
  //       </header>
  //     </div>
  //   );
  // }
}

export default App;
