import React from "react"; 
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Welcome from './src/pages/Welcome'
import Login from './src/pages/Login'
import Signin from './src/pages/Signin'
const App = ()=>{
    return (
        <Router>
            <Routes>
                <Route path="/" element={<Welcome/>} />
                <Route path="/login" element={<Login/>} />
                <Route path="/signup" element={<Signin/>} />
            </Routes>
        </Router>
    )
}
export default App