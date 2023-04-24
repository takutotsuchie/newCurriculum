package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

var updateMutation string = `mutation{
	updateTask(input:
		{explanation: "hey",
		id: "d25124d2-d2b6-9c84-7f9c-56c0a9b71694",
		label_value: 3,
		limit: "2024-05-15T10:30:00Z",
		priority: 4,
		status: Todo,
		title: "heyhey",
		user_id: "3bdb5a00-7ac5-01e4-2b9a-64f787b698db"})
}`
var restoreMutation string = `mutation{
	createTask(input:
		{explanation: "updateTest用",
		id: "d25124d2-d2b6-9c84-7f9c-56c0a9b71694",
		label_value: 3,
		limit: "2024-05-15T10:30:00Z",
		priority: 4,
		status: Todo,
		title: "updateしますわ",
		user_id: "3bdb5a00-7ac5-01e4-2b9a-64f787b698db"})
}`

func TestUpdateTask(t *testing.T) {

	reqBody := MutationRequest{Mutation: updateMutation}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("json.Marshal error", err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal("Post request error", err)
	}
	defer resp.Body.Close()
	var responseStruct UpdateTaskResponse
	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		t.Fatal(err)
	}
	if responseStruct.Data.CreateTask != "d25124d2-d2b6-9c84-7f9c-56c0a9b71694" {
		t.Fatal("invalid ID")
	}
	// 変更を元に戻す
	// ここからupdate2個目

	reqBody = MutationRequest{Mutation: restoreMutation}
	reqBytes, err = json.Marshal(reqBody)
	if err != nil {
		t.Fatal("json.Marshal error", err)
	}

	resp, err = http.Post(url, contentType, bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal("Post request error", err)
	}
	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		t.Fatal(err)
	}
	if responseStruct.Data.CreateTask != "d25124d2-d2b6-9c84-7f9c-56c0a9b71694" {
		t.Fatal("invalid ID")
	}
	t.Log("UpdateTask is OK!")
	defer resp.Body.Close()

}
