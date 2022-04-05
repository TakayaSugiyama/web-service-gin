import React, { useEffect, useState } from 'react';
import '../Member.css';
import axios from 'axios'
import { useParams } from "react-router-dom";

const Member = () => {
  const [member, setMember] = useState({ name: '#', image_link_url: "", nickname: "", birthday_year: 0, birthday_month: 0, birthday_day: 0, blood_type: "" , place_of_birth: "", height: 0, hobby: "", special_skill: "", best_feature: ""})
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

  const birthday_day = (birthday_year: number, birthday_month: number, birthday_day: number) => {
  return `${birthday_year}年${birthday_month}月${birthday_day}日`
}

  useEffect(() => {
    loadData()
  }, [])

  return (
    <div className="Profile">
      <img src={member.image_link_url} alt="photo"/>
      <div className="profile">
        <h1>{ member.name }</h1>
        <dl>
          <dt>ニックネーム</dt>
          <dd>{ member.nickname }</dd>
         </dl>
         <dl>
          <dt>生年月日</dt>
          <dd>{ birthday_day(member.birthday_year, member.birthday_month, member.birthday_day) }</dd>
        </dl>
        <dl>
          <dt>血液型</dt>
          <dd>{ member.blood_type }</dd>
        </dl>
        <dl>
          <dt>出身地</dt>
          <dd>{ member.place_of_birth }</dd>
        </dl>
        <dl>
          <dt>身長</dt>
          <dd>{ `${member.height} cm` }</dd>
        </dl>
         <dl>
          <dt>趣味</dt>
          <dd>{ member.hobby }</dd>
        </dl>
        <dl>
          <dt>チャームポイント</dt>
          <dd>{ member.best_feature}</dd>
        </dl>
       <dl>
          <dt>特技</dt>
          <dd>{ member.special_skill }</dd>
        </dl>
      </div>
    </div>
   )
}

export default Member
