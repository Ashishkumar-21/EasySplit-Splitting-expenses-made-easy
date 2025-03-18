package routes

import (
	"encoding/json"
	"expense-split-backend/models"
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
)

func SettleExpenseHandler(o orm.Ormer, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	var transaction models.Transaction

	err := json.Unmarshal([]byte(request.Body), &transaction)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": allowedOrigin,
			}, Body: "Invalid request body"}, nil
	}
	log.Println("Parsed Request Body:", transaction)

	// Query User
	query := "INSERT INTO global_transactions(PayerID, PayeeID, Amount, Description) values(?,?,?,?)"
	_, errr := o.Raw(query, transaction.PayerID, transaction.PayeeID, transaction.Amount, transaction.Description).Exec()
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
		"message": "settle added",
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
