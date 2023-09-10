import React from 'react';
import ReactDOMClient from 'react-dom/client';


const App = () => {
  return (
    <button >
      Click to get prime numbers
    </button>
  );
};



const root = ReactDOMClient.createRoot(document.getElementById('root'));

root.render(<App />);