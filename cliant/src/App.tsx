import React, { useEffect, useState } from 'react';
import axios from 'axios'
import { Routes, Route, BrowserRouter } from 'react-router-dom'
import Problem from "./components/Problem"
import Member from "./components/Member"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Problem/>}/>
        <Route path="/members/:id" element={<Member/>}/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
