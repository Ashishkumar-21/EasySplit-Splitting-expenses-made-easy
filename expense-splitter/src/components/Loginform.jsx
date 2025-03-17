import { useState } from "react"
import  styles from './Loginform.module.css'
export const Loginform=({isLogin})=>{
    let [ID, setID]=useState("")
    let [MobileNo, setNo]=useState("")
    const handleSubmit= async (e)=>{
        e.preventDefault();//prevents default form submission behavour which reloads the page
        if(!ID && !MobileNo){
            window.alert("please enter ID or MobileNo")
        }
        try{
            const response = await fetch(`https://h1aq3pu22g.execute-api.ap-south-1.amazonaws.com/default/easysplit-login?user_id=${ID}&mobile=${MobileNo}`);
            if(!response.ok){
                throw new Error("invalid credentials")
            }
            console.log(response)
            const data=await response.json()
            const {message, user_id, mobile }= data
            ID=user_id
            MobileNo=mobile
            if (ID) {
                localStorage.setItem("user_id", ID);  
                console.log("user id in ls:", localStorage.getItem("user_id"))
            }
            if (MobileNo) {
                localStorage.setItem("mobile", MobileNo);// Save user_id in localStorage
                console.log("mobile in ls:", localStorage.getItem("mobile"))
            }
            isLogin({ID, MobileNo, message})
            
        }
        catch(error){
            alert(error.message)

        }


    }
    return (
        <div className={styles.loginContainer}>
            <div className={styles.loginBox}>
            <h2 className={styles.loginTitle}> Login</h2>
            <form onSubmit={handleSubmit}>
                <input className={styles.inputField} type="text" placeholder="Enter User ID" value={ID} onChange={(e)=>(setID(e.target.value))}></input>
                <h2>OR</h2>
                <input className={styles.inputField} type="telephone" placeholder="Enter MobileNo" value={MobileNo} onChange={(e)=>(setNo(e.target.value))}></input>
                <button className={styles.loginBtn} type="submit">login</button>
            </form>
            </div>
        </div>
        
    )


}