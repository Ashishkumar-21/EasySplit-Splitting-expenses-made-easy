import React, { useState } from "react";
import  styles from '../components/Loginform.module.css'
import { useNavigate } from "react-router-dom";

export function Addexpense() {
    const [Friend_id, setid]=useState("")
    const [Amount, setcost]=useState("")
    const [Desc, setdesc]=useState("")
    const navigate = useNavigate(); 

    const handleSubmit= async (e)=>{
        e.preventDefault();//prevents default form submission behavour which reloads the page
        if(!Friend_id || !Amount || !Desc){
            window.alert("please enter all the details")
        }
        try{
            const response = await fetch(`https://h1aq3pu22g.execute-api.ap-south-1.amazonaws.com/default/easysplit-login-createExpense`, {
                method: "POST",
                headers: {
                "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    payer_id:localStorage.getItem("user_id"),
                    payee_id:Friend_id,
                    amount:Number(Amount),
                    description:Desc
            }),
        });
        console.log("before invalid Response:", response);
            if(!response.ok){
                throw new Error("invalid credentials")
            }
            
            console.log("before parsing Response:", response);
            let data = await response.json();
            console.log("after parsing Response:", data);
            if(data.message=="expense added"){
                window.alert("Added expense")
                navigate("/dashboard")
            }
        }
        catch(error){
            alert(error.message)
        }

    }
    return (
            <div className={styles.loginContainer}>
                <div className={styles.loginBox}>
                <h2 className={styles.loginTitle}> Add Expense</h2>
                <form onSubmit={handleSubmit}>
                    <input className={styles.inputField} type="text" placeholder="Enter Friend ID" value={Friend_id} onChange={(e)=>(setid(e.target.value))}></input>
                    <input className={styles.inputField} type="number" placeholder="Enter Amount" value={Amount} onChange={(e)=>(setcost(e.target.value))}></input>
                    <input className={styles.inputField} type="text" placeholder="Enter Description" value={Desc} onChange={(e)=>(setdesc(e.target.value))}></input>

                    <button className={styles.loginBtn} type="submit">Submit</button>
                </form>
                </div>
            </div>
            
        )

}