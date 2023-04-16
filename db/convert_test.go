package db

import (
	"fmt"
	"newCurriculum/graph/model"
	"newCurriculum/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

func TestParseTime(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{
			input: "2023-04-15T12:34:56Z",
			want:  time.Date(2023, 4, 15, 12, 34, 56, 0, time.UTC),
		},
		{
			input: "2022-01-01T00:00:00Z",
			want:  time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	fmt.Println("HEloo")
	for _, test := range tests {
		got := ParseTime(test.input)
		if !got.Equal(test.want) {
			t.Errorf("ParseTime(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func Test_shapeInput(t *testing.T) {
	input := model.NewTask{
		ID:          "dda5c15d-9460-e8a0-92aa-5876a629d47c",
		Title:       "programming",
		Explanation: "heyhey",
		Limit:       "2006-01-02T15:04:05Z07:00",
		Priority:    3,
		Status:      model.TaskStatusCompleted,
		UserID:      "e8984c19-33fe-d0ad-5867-a39ae861ec9b",
		LabelValue:  2,
	}

	want := model.Task{
		ID:          "dda5c15d-9460-e8a0-92aa-5876a629d47c",
		Title:       "programming",
		Explanation: "heyhey",
		Limit:       "2006-01-02T15:04:05Z07:00",
		Priority:    3,
		Status:      model.TaskStatusCompleted,
		UserID:      "e8984c19-33fe-d0ad-5867-a39ae861ec9b",
		LabelValue:  2,
	}
	if *shapeInput(input) != want {
		t.Errorf("shapeInput() = %v, want %v", shapeInput(input), want)
	}

}

func Test_generateID(t *testing.T) {
	id := generateID()
	if id == "" {
		t.Error("should not empty")
	}
}

func Test_convertToTask(t *testing.T) {
	newTask := model.NewTask{
		ID:          uuid.New().String(),
		Title:       "Test task",
		Explanation: "This is a test task",
		Limit:       time.Now().Format(time.RFC3339),
		Priority:    1,
		Status:      model.TaskStatusCompleted,
		UserID:      uuid.New().String(),
	}
	expectedTask := models.Task{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Explanation: null.NewString(newTask.Explanation, true),
		Limit:       ParseTime(newTask.Limit),
		Priority:    newTask.Priority,
		Status:      string(newTask.Status),
		UserID:      newTask.UserID,
	}
	output := convertToTask(newTask)
	if output != expectedTask {
		t.Errorf("invalid conversion")
	}

}

func Test_createTaskLabelRealation(t *testing.T) {
	taskID := "f67d98dc-62bb-11ec-8d3d-0242ac130003"
	labelID := "123456"
	relation := createTaskLabelRealation(taskID, labelID)

	// Check the generated ID is not empty.
	if relation.ID == "" {
		t.Error("Expected a non-empty ID, but got an empty string.")
	}

	// Check the task ID and label ID are set correctly.
	if relation.TaskID != taskID {
		t.Errorf("Expected task ID %q, but got %q", taskID, relation.TaskID)
	}
	if relation.LabelID != labelID {
		t.Errorf("Expected label ID %q, but got %q", labelID, relation.LabelID)
	}
}

func Test_checkInput(t *testing.T) {
	input1 := model.NewTask{Title: "test title", Limit: "2024-05-15T10:30:00Z"}
	err := checkInput(input1)
	if err != nil {
		t.Errorf("Error should not have been thrown. Got: %v", err)
	}
	// テストケース2: titleが51文字以上の場合、エラーが発生する
	input2 := model.NewTask{Title: "this title is too long this title is too long this title is too long this title is too long this title is too long this title is too long this title is too long this title is too long this title is too long this title is too long this title is too long", Limit: "2024-05-15T10:30:00Z"}
	err = checkInput(input2)
	if err == nil {
		t.Errorf("Error should have been thrown. Got: %v", err)
	}
}
