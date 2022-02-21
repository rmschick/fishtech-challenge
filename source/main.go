package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)


	var ipAddress = request.Body
	resp, err := http.Get(ipAddress)
	if err != nil {
		log.Fatalln(err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
	}

	sb := string(body)
	log.Printf(sb)

	return events.APIGatewayProxyResponse{Body: sb, StatusCode: 200}, nil
}
func main() {
	lambda.Start(Handler)
}
