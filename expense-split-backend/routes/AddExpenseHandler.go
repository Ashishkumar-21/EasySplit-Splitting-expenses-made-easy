package routes

import (
	"encoding/json"
	"expense-split-backend/models"
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
)

func AddExpenseHandler(o orm.Ormer, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	allowedOrigin := "http://easysplit.com:8080"

	if request.HTTPMethod == "OPTIONS" {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  allowedOrigin,
				"Access-Control-Allow-Methods": "OPTIONS, POST",
				"Access-Control-Allow-Headers": "Content-Type",
			},
			Body: "",
		}, nil
	}
	var Global_transactions models.Global_transactions

	err := json.Unmarshal([]byte(request.Body), &Global_transactions)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": allowedOrigin,
			}, Body: "Invalid request body"}, nil
	}
	log.Println("Parsed Request Body:", Global_transactions)

	// Query User
	query := "INSERT INTO Global_transactions(PayerID, PayeeID, Amount, Description) values(?,?,?,?)"
	_, errr := o.Raw(query, Global_transactions.PayerID, Global_transactions.PayeeID, Global_transactions.Amount, Global_transactions.Description).Exec()
	if errr != nil {
		log.Println(" Database error:", errr)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": allowedOrigin,
			},
			Body: `{"message": "Database error"}`,
		}, nil
	}
	responseBody, _ := json.Marshal(map[string]string{
		"code":    "200",
		"message": "expense added",
	})
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  allowedOrigin,
			"Access-Control-Allow-Methods": "OPTIONS, POST",
			"Access-Control-Allow-Headers": "Content-Type",
		},
		Body: string(responseBody),
	}, nil

}
