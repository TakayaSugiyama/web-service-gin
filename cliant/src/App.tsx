import React, { useEffect } from 'react';
import './App.css';
import axios from 'axios'

function App() {
  useEffect(() => {
    axios.get("http://localhost:8080/randommember")
    .then((res) => {
      console.log(res.data)
    })
    .catch((err) => {
      console.log(err)
    })
  }, [])
  return (
    <div className="App">
      <h1>Hello World</h1>
    </div>
  );
}

export default App;
