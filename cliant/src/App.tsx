import React, { useEffect, useState } from 'react';
import axios from 'axios'
import { Routes, Route, BrowserRouter } from 'react-router-dom'
import Problem from "./components/Problem"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Problem/>} key="problem"/>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
