package db

import (
	"context"
	"database/sql"
	"newCurriculum/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type taskTable struct {
}

func (*taskTable) insert(ctx context.Context, db *sql.DB, task models.Task) error {
	err := task.Insert(ctx, db, boil.Infer())
	return err
}

func (*taskTable) update(ctx context.Context, db *sql.DB, task models.Task) error {
	_, err := task.Update(ctx, db, boil.Infer())
	return err
}

func (*taskTable) delete(ctx context.Context, db *sql.DB, id string) error {
	_, err := models.Tasks(qm.Where("id=?", id)).DeleteAll(ctx, db)
	return err
}
