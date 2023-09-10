import './styles.css'
import React, { useEffect } from 'react'
import ReactDOMClient from 'react-dom/client'

const App = () => {
  const [techniques, setTechniques] = React.useState([])

  useEffect(() => {
    fetch('http://localhost:8080/techniques')
      .then((response) => response.json())
      .then((techniques) => {
        setTechniques(techniques)
      })
  }, [])

  return (
    <>
      <div className="bg-blue-500 text-white p-4">
        <h1 className="text-2xl">Jiu Jitsu Technique Tracker</h1>
      </div>

      <div className="p-4">
        <ul>
          {techniques.map((technique) => (
            <li className="border-b py-2 flex justify-between">
              <span>{technique.Name}</span>
              <div>
                <button className="bg-yellow-300 px-2 py-1 text-sm rounded">
                  Edit
                </button>
                <button className="bg-red-300 px-2 py-1 text-sm rounded">
                  Delete
                </button>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </>
  )
}

const root = ReactDOMClient.createRoot(document.getElementById('root'))

root.render(<App />)
