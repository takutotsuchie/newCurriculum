package table

import (
	"context"
	"database/sql"
	"newCurriculum/infra/boiler"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type TaskTable struct {
}

func (*TaskTable) Insert(ctx context.Context, db *sql.DB, task boiler.Task) error {
	err := task.Insert(ctx, db, boil.Infer())
	return err
}

func (*TaskTable) Update(ctx context.Context, db *sql.DB, task boiler.Task) error {
	_, err := task.Update(ctx, db, boil.Infer())
	return err
}

func (*TaskTable) Delete(ctx context.Context, db *sql.DB, id string) error {
	_, err := boiler.Tasks(qm.Where("id=?", id)).DeleteAll(ctx, db)
	return err
}
