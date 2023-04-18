package e2e

import (
	"testing"
)

// main関数を変更して、DBをtest用のDBにしてからテストを行う。
// hey
func TestCreateUser(t *testing.T) {
	graphqlQuery := `mutation {
        createTask(input:
			{explanation: "programm",
			id: "3bdb5a00-7ac5-01e4-2b9a-64f787b698db",
			label_value: 3,
			limit: "2024-05-15T10:30:00Z",
			priority: 3,
			status: Todo,
			title: "programming",
			user_id: "3bdb5a00-7ac5-01e4-2b9a-64f787b698db"})
    }`
	restoreQuery := `mutation {
		deleteTask(input:
			 {ID: "3bdb5a00-7ac5-01e4-2b9a-64f787b698db"}
			)
	}`

}
