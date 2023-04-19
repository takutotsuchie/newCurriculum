package e2e

type CreateTaskResponse struct {
	Data struct {
		CreateTask string `json:"createTask"`
	} `json:"data"`
}
type DeleteTaskResponse struct {
	Data struct {
		DeleteTask string `json:"deleteTask"`
	} `json:"data"`
}
type UpdateTaskResponse struct {
	Data struct {
		CreateTask string `json:"updateTask"`
	} `json:"data"`
}
type MutationRequest struct {
	Mutation string `json:"query"`
}

const url = "http://localhost:8000/query"
const contentType = "application/json"
