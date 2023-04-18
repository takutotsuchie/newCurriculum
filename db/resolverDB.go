package db

import (
	"context"
	"newCurriculum/db/table"
	"newCurriculum/gql/model"
	"time"

	"github.com/robfig/cron"
)

// 4つのリゾルバはここにつながる。
func CreateTask(ctx context.Context, input model.NewTask) (string, error) {
	db := GetDB()
	var taskTable table.TaskTable
	var labelTable table.LabelTable
	var taskLabelTable table.TaskLabelTable
	err := taskTable.Insert(ctx, db, convertInput(input))
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
	db := GetDB()
	var taskTable table.TaskTable
	var labelTable table.LabelTable
	var taskLabelTable table.TaskLabelTable
	err := taskTable.Update(ctx, db, convertInput(input))
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
func OnLimit(ctx context.Context, userID string) (<-chan string, error) {
	ch := make(chan string, 50)
	go func() {
		c := cron.New()
		c.AddFunc("0 7 * * *", func() {
			taskIDs := searchDB(ctx, userID)
			for _, taskID := range taskIDs {
				ch <- taskID
			}
		})
		c.Start()

		// cronジョブを継続的に実行するために無限ループを使用する
		for {
			time.Sleep(time.Minute)
		}
	}()
	return ch, nil
}
