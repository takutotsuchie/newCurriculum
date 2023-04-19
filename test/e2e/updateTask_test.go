package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestUpdateTask(t *testing.T) {
	type UpdateTaskResponse struct {
		Data struct {
			CreateTask string `json:"updateTask"`
		} `json:"data"`
	}
	updateMutation := `mutation{
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

	reqBody := MutationRequest{Mutation: updateMutation}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Print("json.Marshal error", err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Print("Post request error", err)
	}
	defer resp.Body.Close()
	var responseStruct UpdateTaskResponse
	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(responseStruct)

	// 変更を元に戻す
	// ここからupdate2個目
	restoreMutation := `mutation{
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
	reqBody = MutationRequest{Mutation: restoreMutation}
	reqBytes, err = json.Marshal(reqBody)
	if err != nil {
		log.Print("json.Marshal error", err)
	}

	resp, err = http.Post(url, contentType, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Print("Post request error", err)
	}
	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(responseStruct)
	fmt.Println("UpdateTask is OK!")
	defer resp.Body.Close()

}
