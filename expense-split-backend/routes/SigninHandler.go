package routes

import (
	"encoding/json"
	"expense-split-backend/models"
	"log"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
)

func SigninHandler(o orm.Ormer, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

	var requestBody map[string]string

	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": allowedOrigin,
			}, Body: "Invalid request body"}, nil
	}
	log.Println("Parsed Request Body:", requestBody)

	name := requestBody["Name"]
	mobile := requestBody["mobile"]
	log.Println("request Body:", requestBody)
	log.Println("name :", name, "mobile :", mobile)
	if name == "" || mobile == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  allowedOrigin,
				"Access-Control-Allow-Methods": "OPTIONS, POST",
				"Access-Control-Allow-Headers": "Content-Type",
			}, Body: `{"message": "empty inputs"}`}, nil
	}

	var user models.Userauth

	// Query User
	query := "SELECT * FROM userauth WHERE mobile = ? LIMIT 1"
	error := o.Raw(query, mobile).QueryRow(&user)

	if error == orm.ErrNoRows {
		log.Println(" New user proceeding to create account")
		userID := generateUserID()
		query := "insert into userauth(User_id, Mobile, Name) values(?, ?, ?)"
		_, errr := o.Raw(query, userID, mobile, name).Exec()
		if errr != nil {
			log.Println(" Database error:", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": allowedOrigin,
				},
				Body: `{"message": "Database error:"}`,
			}, nil
		}
		responseBody, _ := json.Marshal(map[string]string{
			"code":      "200",
			"user_id":   userID,
			"user_name": name,
			"mobile":    mobile,
			"message":   "Signin successful",
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

	responseBody, _ := json.Marshal(map[string]string{
		"message": "User already exists please login",
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
