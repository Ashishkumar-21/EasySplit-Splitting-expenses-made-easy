import { useState } from "react";
import { Signinform } from "../components/Signinform";

function Signin() {
  const [user, setUser] = useState(null);

  const handleSignin = (userData) => {
    console.log("User state before update:", userData);
    setUser(userData);
  };

  return (
    <div>
      {user ? (
        user.message === "Signin successful" ? (
          <h1>
            Welcome {user.user_name || "User"}, User ID: {user.user_id || "N/A"}
          </h1>
        ) : (
          <h1>Error: {user.message}</h1>
        )
      ) : (
        <Signinform onSignin={handleSignin} />
      )}
    </div>
  );
}

export default Signin;
