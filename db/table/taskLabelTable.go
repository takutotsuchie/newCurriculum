package table

import (
	"context"
	"database/sql"
	"fmt"
	"newCurriculum/infra/boiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TaskLabelTable struct {
}

func (*TaskLabelTable) Insert(ctx context.Context, db *sql.DB, taskLabelRelation boiler.TaskLabelRelation) error {
	err := taskLabelRelation.Insert(ctx, db, boil.Infer())
	return err
}
func (*TaskLabelTable) Update(ctx context.Context, db *sql.DB, taskLabelRelation boiler.TaskLabelRelation) error {
	_, err := taskLabelRelation.Update(ctx, db, boil.Infer())
	return err
}
func (*TaskLabelTable) Delete(ctx context.Context, db *sql.DB, taskID string) error {
	a, err := boiler.TaskLabelRelations(qm.Where("task_id=?", taskID)).DeleteAll(ctx, db)
	fmt.Println(a)
	return err
}
