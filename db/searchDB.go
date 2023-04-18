package db

import (
	"context"
	"log"
	"newCurriculum/infra/boiler"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// 毎朝7時にそのユーザーのtaskの期限が過ぎているものを探す関数。
func searchDB(ctx context.Context, userID string) []string {
	db := GetDB()
	tasks, err := boiler.Tasks(qm.Where("user_id=?", userID)).All(ctx, db)
	if err != nil {
		log.Print("go routine error")
	}
	var taskIDs []string
	for _, task := range tasks {
		if task.Limit.After(time.Now()) {
			taskIDs = append(taskIDs, task.ID)
		}
	}
	return taskIDs
}
