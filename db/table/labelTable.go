package table

import (
	"context"
	"database/sql"
	"newCurriculum/infra/boiler"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type LabelTable struct {
}

func (*LabelTable) SelectID(ctx context.Context, db *sql.DB, labelValue int) (string, error) {
	ans, err := boiler.Labels(qm.Where("value=?", labelValue)).One(ctx, db)
	return ans.ID, err
}
