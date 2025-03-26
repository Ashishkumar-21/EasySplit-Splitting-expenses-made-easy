# Splitwise
/Splitwise
  /expense-splitter
    /src
      /components
        /Loader.jsx
        /Loader.module.css
        /Loginform.jsx
        /Loginform.module.css
        /Signinform.jsx
      /pages
        /Addexpense.jsx
        /Dashboard.jsx
        /Dashboard.module.css
        /Friend.module.css
        /Friendtransactions.jsx
        /Login.jsx
        /Settleexpense.jsx
        /Signin.jsx
        /Welcome.jsx
        /Welcome.module.css
    /App.jsx
    /index.html
    /main.jsx
  /expense-split-backend
    /events
    /models
      /global_transactions.go
      /userauth.go
    /routes
      /AddExpenseHandler.go
      /Addroutes.go
      /get_friends.go
      /GetFriendHandler.go
      /LoginHandler.go
      /SettleExpenseHandler.go
      /SigninHandler.go
    /main.go

Splitwise Has 2 folders expense-split-backend for backened and expense-splitter for frontend. Expense-splitter has components which stores the components like Loader, Loginform, Signinform. 
Pages consists of welcome, login, signin, dashboard.
App.jsx stores all the routes along with the elements rendered in the frontend.

Expense-split-backend has events which consistes of json test files I used for local sam testing. Models have schema of tables used in the database. I have used the userauth and the Global_transactions table. All endpoints are defined in the Addroutes.go along with their handler function. Main.go has connection of database and migration. 
The environment variable DSN is defined in the aws lambda function 'easysplit-login'.