package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var subscriptionMessage string = `subscription{
	onLimit(input: 
	  {
	  userID : "3bdb5a00-7ac5-01e4-2b9a-64f787b698db",
	  when: now
	  
  })
  }`

type SubscriptionStruct struct {
	Data struct {
		OnLimit string `json:"onLimit"`
	} `json:"data"`
}

func TestSubscription(t *testing.T) {
	reqBody := MutationRequest{Mutation: subscriptionMessage}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post(url, contentType, bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal(err)
	}
	var resStruct SubscriptionStruct
	err = json.NewDecoder(res.Body).Decode(&resStruct)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("subscription is OK!")
}
