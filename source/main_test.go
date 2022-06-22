package main_test

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

const (
	apiURL        = "https://s1cq7a8l4e.execute-api.us-east-1.amazonaws.com/"
	ipAddressPath = "ipaddress"
)

type Client struct {
	restyClient *resty.Client
}

func CreateClient() *Client {
	client := resty.NewWithClient(http.DefaultClient).
		SetHeader("Content-Type", "application/json")

	return &Client{
		restyClient: client,
	}
}

func TestHandler(t *testing.T) {
	tests := []struct {
		name    string
		request string
		expect  string
		err     error
	}{
		{
			name:    "Should return google ip information",
			request: `8.8.8.8`,
			expect:  `{"IPAddress":"8.8.8.8","Hostname":"dns.google","City":"Mountain View","Country":"US","Location":"37.4056,-122.0775","Org":"AS15169 Google LLC"}`,
			err:     nil,
		},
		{
			name:    "Should return frontier ip information",
			request: `47.184.72.33`,
			expect:  `{"IPAddress":"47.184.72.33","Hostname":"47-184-72-33.dlls.tx.frontiernet.net","City":"Dallas","Country":"US","Location":"32.7699,-96.7430","Org":"AS5650 Frontier Communications of America, Inc."}`,
			err:     nil,
		},
		{
			name:    "Should return frontier ip information",
			request: `47.184.72.33`,
			expect:  `{"IPAddress":"47.184.72.33","Hostname":"47-184-72-33.dlls.tx.frontiernet.net","City":"Dallas","Country":"US","Location":"32.7699,-96.7430","Org":"AS5650 Frontier Communications of America, Inc."}`,
			err:     nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var result interface{}

			ctx := context.Background()

			client := CreateClient()

			response, err := client.restyClient.
				SetBaseURL(apiURL).
				R().
				SetContext(ctx).
				SetResult(&result).
				SetBody(test.request).
				Post(ipAddressPath)

			if err != nil {
				log.Fatal(err, response, "bad request")
			}

			assert.Equal(t, test.expect, response.String())
		})

	}
}
