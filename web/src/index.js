import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import Login from './pages/Login'; 
import Home from './pages/Home';  
import Fun1 from './pages/Fun1';
import Fun2 from './pages/Fun2';
import Fun3 from './pages/Fun3';
import Fun4 from './pages/Fun4';


const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />     
        <Route path="/home" element={<Home />} /> 
        <Route path="/Fun1" element={<Fun1 />} /> 
        <Route path="/Fun2" element={<Fun2 />} /> 
        <Route path="/Fun3" element={<Fun3 />} /> 
        <Route path="/Fun4" element={<Fun4 />} /> 
      </Routes>
    </Router>
  </React.StrictMode>
);

reportWebVitals();
