import React, { useEffect, useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import  styles from '../components/Loginform.module.css'

export function FriendTransactions() {
    const nav = useNavigate()
    const location = useLocation();
    const [FriendData, setFriendData] = useState([]);
    const userId = localStorage.getItem("user_id");
    const searchParams = new URLSearchParams(location.search);
    const friendId = searchParams.get("friend_id");

    useEffect(() => {
        if (!userId) {
            alert("User ID not found!");
            return;
        }

        const getFriends = async () => {
            try {
                const response = await fetch(`https://h1aq3pu22g.execute-api.ap-south-1.amazonaws.com/default/easysplit-get_friend_transactions?user_id=${userId}&friend_id=${friendId}`, {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                    },
                });

                if (!response.ok) {
                    throw new Error("Invalid response");
                }

                const data = await response.json();
                console.log("friend Data:", data.data);
                setFriendData(data.data);
            } catch (error) {
                alert(error.message);
            }
        };

        getFriends();
    }, [userId]);

    return (
        <>
            <h1>You and Your friend</h1>
            <button onClick={()=>nav("/addexpense")}>Add expense</button>
            <button onClick={()=>nav("/settleexpense")}>Settle up</button>
            <ul>
                {FriendData ? (
                    FriendData.map((transaction, index) => (
                        <li
                            key={index}
                            className={`${styles.displayelements} ${
                                transaction.amount >= 0
                                    ? styles.positiveBalance
                                    : styles.negativeBalance
                            }`}
                        >
                              {transaction.description} - Amount: {transaction.amount}
                        </li>
                    ))
                ) :<h2>No transactions entered </h2>}
            </ul>
            <button onClick={() => nav("/dashboard")}>Dashboard</button>
        </>
    );
}


