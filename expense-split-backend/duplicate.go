// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-lambda-go/lambda"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var Db *gorm.DB

// type Userauth struct {
// 	User_id string `gorm:"primaryKey" json:"user_id"`
// 	Mobile  string `gorm:"unique" json:"mobile"`
// 	Name    string `json:"name"`
// }


// func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	var inputUser Userauth
// 	errr := json.Unmarshal([]byte(request.Body), &inputUser)
// 	if errr != nil {
// 		log.Println("Error json:", errr)
// 		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "json related error"}, nil
// 	}

// 	log.Println("Received user:", request.Body)
// 	var user Userauth
// 	res := Db.Where("ID=? OR mobilenum=?", inputUser.User_id, inputUser.Mobile).First(&user)

// 	if res != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"messege": "Invalid ID"})
// 	}
// 	ctx.JSON(200, gin.H{"User_id": inputUser.User_id, "mobile": inputUser.Mobile})

// 	return events.APIGatewayProxyResponse{StatusCode: 200, Body: inputUser.Mobile}, nil
// }

// func main() {
// 	ConnectDB()
// 	lambda.Start(handler)
// }
