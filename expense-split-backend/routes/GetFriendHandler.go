package routes

import (
	"encoding/json"
	"expense-split-backend/models"
	"log"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
)

func GetFriendHandler(o orm.Ormer, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	allowedOrigin := "http://easysplit.com:8080"

	FriendID := request.QueryStringParameters["friend_id"]
	UserID := request.QueryStringParameters["user_id"]

	if FriendID == "" || UserID == "" {
		log.Println("friend id or user id is empty")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  allowedOrigin,
				"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type, Authorization",
				"Content-Type":                 "application/json",
			},
			Body: `{"message": "friend id or user id is empty:"}`,
		}, nil
	}

	var results []models.Global_transactions

	query := `SELECT * FROM Global_transactions WHERE (PayerID=? AND PayeeID=?) OR (PayerID=? AND PayeeID=?)`
	numRows, err := o.Raw(query, UserID, FriendID, FriendID, UserID).QueryRows(&results)

	log.Println("Rows fetched:", numRows)

	if err != nil {
		log.Println("Database query error:", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  allowedOrigin,
				"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
				"Access-Control-Allow-Headers": "Content-Type, Authorization",
				"Content-Type":                 "application/json",
			}, Body: `{"message": "Failed to fetch dashboard data"}`}, nil
	}

	responseBody, _ := json.Marshal(map[string]interface{}{
		"code":      200,
		"message":   "Dashboard data retrieved successfully",
		"User ID":   UserID,
		"Friend ID": FriendID,
		"data":      results,
	})
	log.Printf("Query Results: %+v\n", results)
	log.Println("Userid, FriendId", UserID, FriendID)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  allowedOrigin,
			"Access-Control-Allow-Methods": "GET, POST, OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type, Authorization",
			"Content-Type":                 "application/json",
		},
		Body: string(responseBody),
	}, nil
}
