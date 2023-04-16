package db

import (
	"errors"
	"fmt"
	"newCurriculum/graph/model"
	"newCurriculum/models"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

func ParseTime(input string) time.Time {
	time, _ := time.Parse(time.RFC3339, input)
	return time
}

func shapeInput(input model.NewTask) *model.Task {
	return &model.Task{
		ID:          input.ID,
		Title:       input.Title,
		Explanation: input.Explanation,
		Limit:       input.Limit,
		Priority:    input.Priority,
		Status:      input.Status,
		UserID:      input.UserID,
		LabelValue:  input.LabelValue,
	}
}

func convertToTask(newTask model.NewTask) models.Task {
	task := models.Task{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Explanation: null.NewString(newTask.Explanation, newTask.Explanation != ""),
		Limit:       ParseTime(newTask.Limit),
		Priority:    newTask.Priority,
		Status:      string(newTask.Status),
		UserID:      newTask.UserID,
	}
	return task
}
func createTaskLabelRealation(task_id string, label_id string) models.TaskLabelRelation {
	return models.TaskLabelRelation{
		ID:      generateID(),
		TaskID:  task_id,
		LabelID: label_id,
	}
}
func generateID() string {
	return uuid.New().String()
}

//	タスクのタイトルの文字数は 50 文字まで
//
// タスクの説明文の文字数は 300 文字まで
// 終了期限が現在日時以前ではいけない
func checkInput(input model.NewTask) error {
	if len(input.Title) > 50 {
		return errors.New("title長すぎ")
	}
	if len(input.Explanation) > 300 {
		return errors.New("")
	}
	fmt.Println(ParseTime(input.Limit).Before(time.Now()))
	if ParseTime(input.Limit).Before(time.Now()) {
		fmt.Println(input.Limit)
		return fmt.Errorf("終了期限が現在日時以前です%s", input.Limit)
	}
	return nil
}
