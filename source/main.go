package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/ipinfo/go/v2/ipinfo"
)

// checkInputForIp will return true if the user inputs IPv4 dotted decimal ("192.0.2.1"), IPv6 ("2001:db8::68"), or IPv4-mapped IPv6 ("::ffff:192.0.2.1") form and returns false if the user enters a hostname
//this will decide how we handle the information
func checkInputForIp(host string) bool {
	addr := net.ParseIP(host)
	if addr != nil {
		return true
	}
	return false
}

//formatting ip information to send back to client
type IPinformation struct {
	IPAddress string
	Hostname  string
	City      string
	Country   string
	Location  string
	Org       string
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	var userInput = request.Body

	if checkInputForIp(userInput) {
		IPaddress := userInput //get the IP address from userInput

		client := ipinfo.NewClient(nil, nil, "6b0a1ab97c8cb3") //utilize IP information Go Client Library to get IP information https://github.com/ipinfo/go
		info, err := client.GetIPInfo(net.ParseIP(IPaddress))
		if err != nil { //if there was an error in grabbing information then error out
			log.Fatal(err)
			errorNoInformationGrabbed := "Could not get information from IP address\n Please try another IP address or Hostname"

			return events.APIGatewayProxyResponse{Body: errorNoInformationGrabbed, StatusCode: 400}, nil
		}

		informationToPass, err := json.Marshal(IPinformation{IPAddress: IPaddress, Hostname: info.Hostname, City: info.City, Country: info.Country, Location: info.Location, Org: info.Org}) //make a json of the IP information struct to pass along to API Gateway
		if err != nil {
			panic(err)
		}

		return events.APIGatewayProxyResponse{Body: string(informationToPass), StatusCode: 200}, nil //all went smoothly, send the user all the information
	} else { //if the user sent a HOSTNAME
		hostName := userInput           //get Hostname from user
		resp, err := http.Get(hostName) //send a http request to hostname
		if err != nil {                 //if there was an error in getting the http response, send error
			log.Fatalln(err)
			errorNotValidInput := "Error: not valid input or could not grab information from domain. \nPlease utilize IPv4, IPv6 or a valid url to get information."
			return events.APIGatewayProxyResponse{Body: errorNotValidInput, StatusCode: 400}, nil
		}

		body, err := ioutil.ReadAll(resp.Body) //read the body of the http response
		if err != nil {                        //if there was an error, send error to user
			log.Fatalln(err)
			errorNothingToRead := "Error: could not read information from url"
			return events.APIGatewayProxyResponse{Body: errorNothingToRead, StatusCode: 400}, nil
		}

		sb := string(body) //convert json to string
		log.Printf(sb)

		return events.APIGatewayProxyResponse{Body: sb, StatusCode: 200}, nil //send response to user
	}
}

func main() {
	lambda.Start(Handler)
}
