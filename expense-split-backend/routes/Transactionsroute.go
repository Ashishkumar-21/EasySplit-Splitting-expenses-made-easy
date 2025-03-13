package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
)

// DashboardResponse struct to hold outstanding transactions
type DashboardResponse struct {
	FriendID   string  `json:"friend_id"`
	Name       string  `json:"name"`
	NetBalance float64 `json:"netbalance"`
}

// GetDashboardHandler retrieves the dashboard data for a user
func GetDashboardHandler(o orm.Ormer, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	allowedOrigin := "http://easysplit.com:8080"

	userID := request.QueryStringParameters["user_id"]
	if userID == "" {
		return events.APIGatewayProxyResponse{StatusCode: http.StatusBadRequest, Headers: map[string]string{
			"Access-Control-Allow-Origin": allowedOrigin,
		}, Body: `{"message": "User ID is required"}`}, nil
	}

	var results []DashboardResponse

	query := `
    WITH balances AS (
        SELECT 
            owed_by AS friend_id,
            SUM(cost / 2) AS amount
        FROM global_transactions
        WHERE paid_by = ?
        GROUP BY owed_by
        
        UNION ALL
        
        SELECT 
            paid_by AS friend_id,
            -SUM(cost / 2) AS amount
        FROM global_transactions
        WHERE owed_by = ?
        GROUP BY paid_by
    )
    SELECT 
        g.friend_id,
        u.name,
        SUM(g.amount) AS netbalance
    FROM balances g
    JOIN global_users u ON g.friend_id = u.user_id
    GROUP BY g.friend_id, u.name;
    `

	_, err := o.Raw(query, userID, userID).QueryRows(&results)
	if err != nil {
		log.Println("Database query error:", err)
		return events.APIGatewayProxyResponse{StatusCode: http.StatusInternalServerError, Headers: map[string]string{
			"Access-Control-Allow-Origin": allowedOrigin,
		}, Body: `{"message": "Failed to fetch dashboard data"}`}, nil
	}

	responseBody, _ := json.Marshal(map[string]interface{}{
		"code":    "200",
		"message": "Dashboard data retrieved successfully",
		"data":    results,
	})

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": allowedOrigin,
		},
		Body: string(responseBody),
	}, nil
}
