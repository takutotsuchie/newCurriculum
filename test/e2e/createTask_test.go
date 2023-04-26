package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// main関数を変更して、DBをtest用のDBにしてからテストを行う。
// hey

var createMutation string = `mutation{
	createTask(input:
		{explanation: "programm",
		id: "9995f552-6549-05d3-c60e-2fce4d84e2ff",
		label_value: 3,
		limit: "2024-05-15T10:30:00Z",
		priority: 3,
		status: Todo,
		title: "programming",
		user_id: "3bdb5a00-7ac5-01e4-2b9a-64f787b698db"})
}`
var deleteMutation string = ` mutation{
	deleteTask(input:
		 "9995f552-6549-05d3-c60e-2fce4d84e2ff"
		)
}`

func TestCreateTaskAndDeleteTask(t *testing.T) {
	// リクエストをjsonに変換する
	reqBody := MutationRequest{Mutation: createMutation}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal("json.Marshal error", err)
	}

	resp, err := http.Post(url, contentType, bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal("Post request error", err)
	}
	defer resp.Body.Close()
	// レスポンスを読み込んで処理する
	var responseStruct CreateTaskResponse

	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		t.Fatal(err)
	}
	if responseStruct.Data.CreateTask != "9995f552-6549-05d3-c60e-2fce4d84e2ff" {
		t.Fatal("invalid ID")
	}
	t.Log("CreateTask is OK!")
	// ここからdelete

	reqBody = MutationRequest{Mutation: deleteMutation}
	reqBytes, err = json.Marshal(reqBody)
	if err != nil {
		t.Fatal("json.Marshal error", err)
	}
	resp, err = http.Post(url, contentType, bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Fatal("post error", err)
	}

	var DeleteResponseStruct DeleteTaskResponse

	err = json.NewDecoder(resp.Body).Decode(&DeleteResponseStruct)
	if err != nil {
		t.Fatal(err)
	}
	if DeleteResponseStruct.Data.DeleteTask != "9995f552-6549-05d3-c60e-2fce4d84e2ff" {
		t.Fatal("invalid ID")
	}
	t.Log("DeleteTask is OK!")
}
