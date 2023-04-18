package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestUserQuery(t *testing.T) {
	// まずクエリをかく。
	query := `
        {
            user(id: "123") {
                id
                name
            } 
        }
    `
	// GraphQLリクエストを定義する
	reqBody := map[string]string{
		"query": query,
	}
	// json形式にする。
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(reqBytes))

	// GraphQLサーバーにPOSTリクエストを送信する
	resp, err := http.Post("http://localhost:8000/query", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	type Response struct {
		Data struct {
			User struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		} `json:"data"`
	}
	// レスポンスを読み込んで処理する
	var respBody Response
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", respBody)
	fmt.Println("俺の名前は", respBody.Data.User.Name)
}
