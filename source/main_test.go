// package main_test

// import (
// 	"testing"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/stretchr/testify/assert"
// )

// func TestHandler(t *testing.T){
// 	tests := []struct {
// 		request events.APIGatewayProxyRequest
// 		expect  string
// 		err     error
// 	}{
// 		{
// 			request: events.APIGatewayProxyRequest{Body: "8.8.8.8"},
// 			expect: `{"IPAddress":"8.8.8.8","Hostname":"dns.google","City":"Mountain View","Country":"US","Location":"37.4056,-122.0775","Org":"AS15169 Google LLC"}`,
// 			err:	nil,
// 		},
// 		{
// 			request: events.APIGatewayProxyRequest{Body: "47.184.72.33"},
// 			expect:  `{"IPAddress":"47.184.72.33","Hostname":"47-184-72-33.dlls.tx.frontiernet.net","City":"Denton","Country":"US","Location":"33.1428,-97.0727","Org":"AS5650 Frontier Communications of America, Inc."}`,
// 			err:     nil,
// 		},
// }

// 	for _, test := range tests {
// 		response, err := main.Handler(test.request)
// 		assert.IsType(t, test.err, err)
// 		assert.Equal(t, test.expect, response.Body)
// 	}
// }