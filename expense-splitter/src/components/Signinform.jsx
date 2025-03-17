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
            
            console.log("before parsing Response:", response);
            let data = await response.json();
            console.log("after parsing Response:", data);
            if(data.user_id){
                localStorage.setItem("user_id", data.user_id)
                console.log("user_id ls", localStorage.getItem("user_id"))
            }
            if(data.mobile){
                localStorage.setItem("mobile", data.mobile)
                console.log("mobile ls", localStorage.getItem("mobile"))
            }

            onSignin(data); // Pass the correct object
            
        }
        catch(error){
            alert(error.message)
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