import { useState } from "react";
import { Loginform } from "../components/Loginform";
import { useNavigate } from "react-router-dom";


function Login(){
  const [user, setUser] = useState(null)
  const navigate = useNavigate(); 

  const handleLogin=(userData)=>{
    setUser(userData)
    if (userData.message === "Login successful") {
      navigate("/dashboard")} 
      
  }
  return(
    <div>
      {user?(<h1>Welcome {user.ID}</h1>)
      :(<Loginform isLogin={handleLogin}/>)}

    </div>
  )
}
export default Login