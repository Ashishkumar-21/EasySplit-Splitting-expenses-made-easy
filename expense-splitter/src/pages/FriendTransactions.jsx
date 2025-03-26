import React, { useEffect, useState } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import styles from "./Friend.module.css";
import { Loader } from "../components/Loader";


export function FriendTransactions() {
    const nav = useNavigate();
    const location = useLocation();
    const [FriendData, setFriendData] = useState([]);
    const userId = localStorage.getItem("user_id");
    const searchParams = new URLSearchParams(location.search);
    const friendId = searchParams.get("friend_id");
    const [FriendName, setFriendName] = useState("");
    const [Balance, setBalance] = useState(0);
    const [loading, setLoading] = useState(false)

    useEffect(() => {
        if (!userId) {
            alert("User ID not found!");
            return;
        }

        const getFriends = async () => {
            setLoading(true);
            try {
                const response = await fetch(
                    `https://h1aq3pu22g.execute-api.ap-south-1.amazonaws.com/default/easysplit-get_friend_transactions?user_id=${userId}&friend_id=${friendId}`,
                    {
                        method: "GET",
                        headers: {
                            "Content-Type": "application/json",
                        },
                    }
                );

                if (!response.ok) {
                    throw new Error("Invalid response");
                }

                const data = await response.json();
                setFriendName(data.Friend_Name);
                setFriendData(data.data);
                setBalance(data.Balance);
            } catch (error) {
                alert(error.message);
            } finally {
                setLoading(false)
            }
        };

        getFriends();
    }, [userId]);

    return (
        <>
            {loading ? <Loader /> :
                <div className={styles.body}>
                    <div className={styles.friendTransactionsContainer}>
                        <div className={styles.dashboardHeader}>
                            <h1 className={styles.friendHeading}>You and {FriendName}</h1>
                            <div className={styles.buttonContainer}>
                                <button onClick={() => nav(`/addexpense?&friend_id=${friendId}`)}>
                                    Add
                                </button>
                                <button onClick={() => nav(`/settleexpense?&friend_id=${friendId}`)}>
                                    Settle
                                </button>
                                <button onClick={() => nav("/dashboard")}>Dashboard</button>
                            </div>
                        </div>
                        <h2 className={styles.balanceInfo}>
                            {Balance > 0 ? `You get ₹${Math.abs(Balance)}` : `You owe ₹${Math.abs(Balance)}`}
                        </h2>
                        <ul className={styles.transactionList}>
                            {FriendData? (
                                FriendData.map((transaction, index) => (
                                    <li
                                        key={index}
                                        className={`${styles.transactionItem} ${transaction.status === "Paid"
                                            ? styles.positiveBalance
                                            : styles.negativeBalance
                                            }`}
                                    >
                                        {transaction.description === "settle"
                                            ? `${transaction.description} - Amount: ₹${transaction.amount}`
                                            : `${transaction.description} - Amount: ₹${transaction.amount} - Share: ₹${transaction.share} (${transaction.status})`}
                                    </li>
                                ))
                            ) : (
                                <h2>No transactions entered</h2>
                            )}
                        </ul>

                    </div>

                </div>}
        </>


    );
}
