package db

import (
	"errors"
	"fmt"
	"newCurriculum/gql/model"
	"newCurriculum/infra/boiler"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

func ParseTime(input string) time.Time {
	time, _ := time.Parse(time.RFC3339, input)
	return time
}

func createTaskLabelRealation(task_id string, label_id string) boiler.TaskLabelRelation {
	return boiler.TaskLabelRelation{
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

func convertInput(input model.NewTask) boiler.Task {
	return boiler.Task{
		ID:          input.ID,
		Title:       input.Title,
		Explanation: null.NewString(input.Explanation, true),
		Limit:       ParseTime(input.Limit),
		Priority:    input.Priority,
		Status:      input.Status.String(),
		UserID:      input.UserID,
	}
}
