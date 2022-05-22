import React from 'react'
import './App.css'
import Nav from './components/Navigation';
import MainRoutes from './Routes';

function App() {
  return (
    <div className="App">
      <Nav />
      <MainRoutes />
    </div>
   
  );
}

export default App;
