package routes

import (
	"encoding/json"
	"expense-split-backend/models"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
)

func LoginHandler(o orm.Ormer, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ID := request.QueryStringParameters["user_id"]
	MobileNo := request.QueryStringParameters["mobile"]

	log.Println("ID:", ID)
	log.Println("MobileNo:", MobileNo)
	log.Println("Raw Event:", request)

	var user models.Userauth

	// Query User
	query := "SELECT * FROM userauth WHERE user_id = ? OR mobile = ? LIMIT 1"
	err := o.Raw(query, ID, MobileNo).QueryRow(&user)

	if err != nil {
		if err == orm.ErrNoRows {
			log.Println(" User not found")
			return events.APIGatewayProxyResponse{StatusCode: 400, Headers: map[string]string{
				"Access-Control-Allow-Origin": "*"}, Body: "Invalid credentials"}, nil
		}
		log.Println(" Database error:", err)
		return events.APIGatewayProxyResponse{StatusCode: 500, Headers: map[string]string{
			"Access-Control-Allow-Origin": "*"}, Body: fmt.Sprintf("Database error: %s", err)}, nil
	}

	responseBody, _ := json.Marshal(map[string]string{
		"code":      "200",
		"user_id":   user.User_id,
		"user_name": user.Name,
		"mobile":    user.Mobile,
		"message":   "Login successful",
	})
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type",
		},
		Body: string(responseBody),
	}, nil
}
