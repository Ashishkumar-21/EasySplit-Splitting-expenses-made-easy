import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
// import './Dashboard.module.css'
export function Dashboard() {
    const nav = useNavigate()
    const [dashboardData, setDashboardData] = useState([]);
    const userId = localStorage.getItem("user_id");
    const Mobile = localStorage.getItem("mobile");
    const [userName, setName] = useState(0);
    const [netbalance, setNet]=useState(0);
    const [positivebalance, setPos]=useState(0);

    useEffect(() => {
        if (!userId) {
            alert("User ID not found!");
            return;
        }

        const getFriends = async () => {
            try {
                const response = await fetch(`https://h1aq3pu22g.execute-api.ap-south-1.amazonaws.com/default/easysplit-get_friends?user_id=${userId}&mobile=${Mobile}`, {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json",
                    },
                });

                if (!response.ok) {
                    throw new Error("Invalid response");
                }

                const data = await response.json();
                console.log("Dashboard Data:", data);
                setDashboardData(data.data);
                setNet(data.Balance)
                setPos(data.PositiveBalance)
                setName(data.userName);
            } catch (error) {
                alert(error.message);
            }
        };

        getFriends();
    }, [userId]);

    return (
        <div className="body">
            <h1>Dashboard</h1>
            <h2>Hello {userName}!!</h2>

            <div className="component">
                <div className="Balances">
                    Balance: {netbalance}
                </div>
                <div className="Balances">
                    You Owe: {netbalance-positivebalance}
                </div>
                <div className="Balances">
                    You are Owed : {positivebalance}
                </div>
            </div>
            <button onClick={()=>nav("/addexpense")}>Add expense</button>
            <button onClick={()=>nav("/settleexpense")}>Settle up</button>
            
            <ul>
                {dashboardData?dashboardData.map((friend, index) => (
                    <li onClick={()=>nav(`/friend-transactions?friend_id=${friend.friend_id}`)} key={index} className={`displayelements ${friend.netbalance >= 0 ? "positiveBalance" : "negativeBalance"}`}>
                        {friend.name} (ID: {friend.friend_id}) - Balance: {friend.netbalance}

                    </li>
                )):<h3>No transactions entered </h3>}
            </ul>
        </div>
    );
}


