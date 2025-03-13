import { useState } from "react";
import { Loginform } from "../components/Loginform";

function Login(){
  const [user, setUser] = useState(null)

  const handleLogin=(userData)=>{
    setUser(userData)
  
  }
  return(
    <div>
      
      {user?(<h1>Welcome {user.user_name}</h1>)
      :(<Loginform isLogin={handleLogin}/>)}

    </div>
  )
}
export default Login