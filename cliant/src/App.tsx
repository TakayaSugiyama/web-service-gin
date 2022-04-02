import React, { useEffect, useState } from 'react';
import './App.css';
import axios from 'axios'

function App() {
  const [options, setOptions] = useState([])
  const [member, setMember] = useState({image_link_url: "#"})

  useEffect(() => {
    axios.get("http://localhost:8080/randommember")
    .then((res) => {
      setOptions(res.data.options)
      setMember(res.data.member)
    })
    .catch((err) => {
      console.log(err)
    })
  }, [])
  return (
    <div className="App">
      <h1 className="questin-title">Who is She?</h1>
      <div className="question">
        <img src={member.image_link_url} alt="photo"/>
      </div>
      <div className="memberOptions">
        {
          options.map(elm => <div className="memberOption">{ elm }</div>)
        }
      </div>
    </div>
  );
}

export default App;
