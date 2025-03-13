import { useState } from "react"
import  styles from './Loginform.module.css'
export const Loginform=({isLogin})=>{
    const [ID, setID]=useState("")
    const [MobileNo, setNo]=useState("")
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
            const {user_name}= data
            isLogin({ID, MobileNo, user_name})
        }
        catch(error){
            alert(error.messege)

        }


    }
    return (
        <div className={styles.loginContainer}>
            <div className={styles.loginBox}>
            <h2 className={styles.loginTitle}> Login</h2>
            <form onSubmit={handleSubmit}>
                <input className={styles.inputField} type="number" placeholder="Enter User ID" value={ID} onChange={(e)=>(setID(e.target.value))}></input>
                <h2>OR</h2>
                <input className={styles.inputField} type="telephone" placeholder="Enter MobileNo" value={MobileNo} onChange={(e)=>(setNo(e.target.value))}></input>
                <button className={styles.loginBtn}>login</button>
            </form>
            </div>
        </div>
        
    )


}