import { useState } from "react"
import  styles from './Loginform.module.css'
export const Signinform=({onSignin})=>{
    const [Name, setName]=useState("")
    const [MobileNo, setNo]=useState("")
    const handleSubmit= async (e)=>{
        e.preventDefault();//prevents default form submission behavour which reloads the page
        if(!Name || !MobileNo){
            window.alert("please enter Name and MobileNo")
        }
        try{
            const response = await fetch(`https://h1aq3pu22g.execute-api.ap-south-1.amazonaws.com/default/easysplit-signin`, {
                method: "POST",
                headers: {
                "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    Name: Name, // Send the user input as JSON
                    mobile: MobileNo,
            }),
        });
            if(!response.ok){
                throw new Error("invalid credentials")
            }
            
            console.log(response)
            let data = await response.json();
            console.log("Raw API Response:", data);
            
            if (typeof data === "string") {
                try {
                    data = JSON.parse(data);
                } catch (err) {
                    console.error("Error parsing JSON:", err);
                }
            }
            
            if (data.body) {
                try {
                    data.body = JSON.parse(data.body);
                } catch (err) {
                    console.error("Error parsing body JSON:", err);
                }
            }
            
            console.log("Parsed API Response:", data);
            onSignin(data.body); // Pass the correct object
            
            
            onSignin(u)
        }
        catch(error){
            alert(error.messege)
        }

    }
    return (
        <div className={styles.loginContainer}>
            <div className={styles.loginBox}>
            <h2 className={styles.loginTitle}> Signin</h2>
            <form onSubmit={handleSubmit}>
                <input className={styles.inputField} type="text" placeholder="Enter User Name" value={Name} onChange={(e)=>(setName(e.target.value))}></input>
                <input className={styles.inputField} type="text" placeholder="Enter MobileNo" value={MobileNo} onChange={(e)=>(setNo(e.target.value))}></input>
                <button className={styles.loginBtn} type="submit">Signin</button>
            </form>
            </div>
        </div>
        
    )


}