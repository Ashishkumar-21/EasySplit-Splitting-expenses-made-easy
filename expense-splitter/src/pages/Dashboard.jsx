import React, { useEffect, useState } from "react";

export function Dashboard() {
    const [dashboardData, setDashboardData] = useState([]);
    const userId = localStorage.getItem("user_id");
    const Mobile = localStorage.getItem("mobile");

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
            } catch (error) {
                alert(error.message);
            }
        };

        getFriends();
    }, [userId]);

    return (
        <>
            <h1>Dashboard</h1>
            <button>Add expense</button>
            <button>Settle up</button>
            <ul>
                {dashboardData?dashboardData.map((friend, index) => (
                    <li key={index}>
                        {friend.name} (ID: {friend.friend_id}) - Balance: {friend.netbalance}
                    </li>
                )):<h2>No transactions entered </h2>}
            </ul>
        </>
    );
}


