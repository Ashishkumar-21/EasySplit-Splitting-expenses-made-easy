package main

import (
	"expense-split-backend/models"
	"expense-split-backend/routes"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectDB() {
	// Register MySQL database
	dsn := os.Getenv("DSN")
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		log.Fatal("Failed to register database:", err)
		return
	}

	// Register models
	orm.RegisterModel(new(models.Userauth))

	// Sync database (auto-create tables)
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal("Database migration failed:", err)
		return
	}

	fmt.Println("Database connected and migrated successfully!")
}

func main() {
	ConnectDB()

	// Initialize Lambda routes
	routes.InitLambda()

	// CORS middleware setup

	// log.Fatal(http.ListenAndServe(":9000", handler))
}
