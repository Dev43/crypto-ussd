
import React from 'react';
import { render } from 'react-dom';
import CryptoUSSD from './pages/index';

class App extends React.Component {
  render() {
    return <CryptoUSSD />;
  }
}

render(<App />, document.getElementById('app'));
