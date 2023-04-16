package db

import (
	"context"
	"database/sql"
	"newCurriculum/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type labelTable struct {
}

func (*labelTable) selectID(ctx context.Context, db *sql.DB, labelValue int) (string, error) {
	ans, err := models.Labels(qm.Where("value=?", labelValue)).One(ctx, db)
	return ans.ID, err
}
