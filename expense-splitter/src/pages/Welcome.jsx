
import { useNavigate } from "react-router-dom";
export default function Welcome() {
	const nav = useNavigate()
	
	return(
		<div>
		<h1>Welcome to easy split</h1>
		<button onClick={()=>nav("/login")}>Login</button>
		<button onClick={()=>nav("/signup")}>Signup</button>
		
        
		</div>
	);
}
