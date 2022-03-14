import React, { useEffect, useState } from 'react';
import axios from 'axios'
import "./App.css"

interface MemberInfo {
  id: number,
  name: string,
  profile_link: string,
  team_name: string,
  image_link_url: string,
}

function App() {
  const [members, setMemebers] = useState<MemberInfo | any>([])
  useEffect(() => {
    axios.get("http://localhost:8080/randommembers")
    .then((res) => {
      setMemebers(() => res.data)
    })
    .catch((err) => {
      console.log(err)
    })
  }, [])

  return (
    <>
      { members.map((member: MemberInfo) => {
        return (
          <div className={ "App" }>
            <img src={ member.image_link_url } alt={"photo"}/>
          </div>
        )
      })}
    </>
  );
}

export default App;
