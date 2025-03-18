// package main

// import (
// 	"expense-split-backend/models"
// 	"expense-split-backend/routes"

// 	"fmt"
// 	"log"

// 	// "github.com/gin-gonic/gin"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var Db *gorm.DB

// func ConnectDB() {
// 	dsn := "root:Vaishuveera@2@tcp(0.tcp.in.ngrok.io:12200)/Splitwiser?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to the database:", err)
// 		return
// 	}
// 	Db = db
// 	fmt.Println("Database connected successfully!")

// 	err = Db.AutoMigrate(&models.Userauth{})
// 	if err != nil {
// 		log.Fatal("var migration Failed")
// 		return
// 	}
// 	fmt.Println("Migration sucess")
// }

// func main() {
// 	ConnectDB()
// 	// router := gin.Default()
// 	// router.Run(":8088")
// 	routes.InitLambda()

// }

// https://h1aq3pu22g.execute-api.ap-south-1.amazonaws.com/default/easysplit-login?user_id=12345&mobile=9876543210

package main

import (
	"expense-split-backend/models"
	"expense-split-backend/routes"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func ConnectDB() {
	// Register MySQL database
	err := orm.RegisterDataBase("default", "mysql", "root:Vaishuveera@2@tcp(0.tcp.in.ngrok.io:12615)/Splitwiser?charset=utf8&parseTime=True&loc=Local")
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
	routes.InitLambda() // Start Lambda function
}
