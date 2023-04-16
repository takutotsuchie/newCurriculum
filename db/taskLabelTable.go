package db

import (
	"context"
	"database/sql"
	"fmt"
	"newCurriculum/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type taskLabelTable struct {
}

func (*taskLabelTable) insert(ctx context.Context, db *sql.DB, taskLabelRelation models.TaskLabelRelation) error {
	err := taskLabelRelation.Insert(ctx, db, boil.Infer())
	return err
}
func (*taskLabelTable) update(ctx context.Context, db *sql.DB, taskLabelRelation models.TaskLabelRelation) error {
	_, err := taskLabelRelation.Update(ctx, db, boil.Infer())
	return err
}
func (*taskLabelTable) delete(ctx context.Context, db *sql.DB, taskID string) error {
	a, err := models.TaskLabelRelations(qm.Where("task_id=?", taskID)).DeleteAll(ctx, db)
	fmt.Println(a)
	return err
}
