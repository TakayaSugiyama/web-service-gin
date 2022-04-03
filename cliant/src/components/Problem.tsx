import React, { useEffect, useState } from 'react';
import '../Problem.css';
import axios from 'axios'

const Problem = () => {
  const [options, setOptions] = useState([])
  const [member, setMember] = useState({image_link_url: "#", name: "", id: 0})
  const [correctCount, setCorrectCount] = useState(0)


  const loadData = async(prevID?:number) => {
    let url: string = ""
    if(prevID === undefined){
      url = `http://localhost:8080/randommember`
    }else{
      url = `http://localhost:8080/randommember?prevID=${prevID}`
    }
    axios.get(url)
    .then((res) => {
      setOptions(res.data.options)
      setMember(res.data.member)
    })
    .catch((err) => {
      console.log(err)
    })
  }

  const checkAnswer = (elm: string):void => {
    if(elm === member.name){
      alert("正解です")
      setCorrectCount(correctCount+1)
    }else{
      alert(`不正解です。正解は${member.name}でした。`)
      setCorrectCount(0)
    }
    loadData(member.id)
  }

  useEffect(() => {
    loadData()
  }, [])

  return (
    <div className="App">
      <h1 className="questin-title">Who is She?</h1>
      <h2 className="answer-count">{ correctCount }問連続正解中</h2>
      <div className="question">
        <img src={member.image_link_url} alt="photo"/>
      </div>
      <div className="memberOptions">
        {
        options.map(elm => <div className="memberOption" onClick={() => checkAnswer(elm)} key={elm}>{ elm }</div>)
        }
      </div>
    </div>
   )
}

export default Problem
