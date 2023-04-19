package db

import (
	"context"
	"fmt"
	"newCurriculum/db/table"
	"newCurriculum/gql/model"
	"time"
)

// 4つのリゾルバはここにつながる。
func CreateTask(ctx context.Context, input model.NewTask) (string, error) {
	err := checkInput(input)
	if err != nil {
		return "", err
	}
	db := GetDB()
	var taskTable table.TaskTable
	var labelTable table.LabelTable
	var taskLabelTable table.TaskLabelTable
	err = taskTable.Insert(ctx, db, convertInput(input))
	if err != nil {
		return "", err
	}
	labelID, err := labelTable.SelectID(ctx, db, input.LabelValue)
	if err != nil {
		return "", err
	}
	err = taskLabelTable.Insert(ctx, db, createTaskLabelRealation(input.ID, labelID))
	if err != nil {
		return "", err
	}
	return input.ID, nil
}
func UpdateTask(ctx context.Context, input model.NewTask) (string, error) {
	err := checkInput(input)
	if err != nil {
		return "", err
	}
	db := GetDB()
	var taskTable table.TaskTable
	var labelTable table.LabelTable
	var taskLabelTable table.TaskLabelTable
	err = taskTable.Update(ctx, db, convertInput(input))
	if err != nil {
		return "", err
	}
	labelID, err := labelTable.SelectID(ctx, db, input.LabelValue)
	if err != nil {
		return "", err
	}
	err = taskLabelTable.Update(ctx, db, createTaskLabelRealation(input.ID, labelID))
	if err != nil {
		return "", err
	}
	return input.ID, nil
}
func DeleteTask(ctx context.Context, input string) (string, error) {
	db := GetDB()
	var taskTable table.TaskTable
	var taskLabelTable table.TaskLabelTable
	err := taskLabelTable.Delete(ctx, db, input)
	if err != nil {
		return "", err
	}
	err = taskTable.Delete(ctx, db, input)
	if err != nil {
		return "", err
	}
	return input, nil
}

// 期限が来たら、通知する。
func OnLimit(ctx context.Context, input model.Limit) (<-chan string, error) {
	ch := make(chan string, 50)
	if input.When == model.WhenTypeDefault {
		go func() {
			// 7時までの時間を計算
			t := time.Now()
			target := time.Date(t.Year(), t.Month(), t.Day()+1, 7, 0, 0, 0, time.Local)
			duration := target.Sub(t)
			// 7時まで、待機
			fmt.Printf("Waiting for %v\n", duration)
			time.Sleep(duration)
			taskIDs := searchDB(ctx, input.UserID)
			for _, taskID := range taskIDs {
				ch <- taskID
			}
		}()
	} else {
		// テスト用のコード
		go func() {
			taskIDs := searchDB(ctx, input.UserID)
			for _, taskID := range taskIDs {
				ch <- taskID
			}
		}()
	}

	return ch, nil
}
