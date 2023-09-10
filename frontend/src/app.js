import './styles.css';
import React from 'react';
import ReactDOMClient from 'react-dom/client';


const App = () => {
  return (
    <button className="bg-blue-400" >
      Click to get prime numbers
    </button>
  );
};

const root = ReactDOMClient.createRoot(document.getElementById('root'));

root.render(<App />);