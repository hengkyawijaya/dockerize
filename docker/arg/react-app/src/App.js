import React from 'react';
import logo from './logo.svg';
import './App.css';
import config from './config';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          {config.SERVICE_NAME}
        </p>
        <a
          className="App-link"
          href={config.LINK}
          target="_blank"
          rel="noopener noreferrer"
        >
          {config.SERVICE_NAME}
        </a>
      </header>
    </div>
  );
}

export default App;
