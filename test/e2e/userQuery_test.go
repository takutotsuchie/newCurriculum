package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

type GraphQLRequestBody struct {
	Query string `json:"query"`
}

func TestUserQuery(t *testing.T) {
	// まずクエリをかく。
	// これはまだjson形式でないので、パースできない
	// graphQLのクエリと、jsonは異なる。
	query := `
        {
            user(id: "123") {
                id
                name
            } 
        }
    `
	// GraphQLリクエストを定義する
	// "query"がkeyで、queryがvalueの一対のjson、雑なjson
	b := GraphQLRequestBody{Query: query}
	// json形式にする。
	reqBytes, err := json.Marshal(b)
	if err != nil {
		t.Error(err)
	}

	// GraphQLサーバーにPOSTリクエストを送信する
	resp, err := http.Post("http://localhost:8000/query", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	// 入れ子の構造体を作る。
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
		t.Error(err)
	}
	t.Logf("%+v\n", respBody)
}
