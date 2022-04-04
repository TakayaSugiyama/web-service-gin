import React, { useEffect, useState } from 'react';
import '../Member.css';
import axios from 'axios'
import { useParams } from "react-router-dom";

const Member = () => {
  const [member, setMember] = useState({ name: '#' })
  const { id }= useParams()


  const loadData = async() => {
    const url = `http://localhost:8080/memberdetails/${id}`
    axios.get(url)
    .then((res) => {
      setMember(res.data)
    })
    .catch((err) => {
      console.log(err)
    })
  }

  useEffect(() => {
    loadData()
  }, [])

  return (
    <div className="App">
      <h1>{ member.name }</h1>
    </div>
   )
}

export default Member
