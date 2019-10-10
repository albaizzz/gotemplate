package sqldb

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//ExecuteInTx executes transaction in stub
//args:
//	ctx: context.Context
//	tx: sql transaction: *sqlx.Tx
//	efunc: stub wrapped inside transaction, will be rollbacked if it fails
//returns:
//	err: operation error
func ExecuteInTx(ctx context.Context, tx *sqlx.Tx, efunc func() error) (err error) {
	err = efunc()
	if err == nil {
		err = tx.Commit()
	} else {
		_ = tx.Rollback()
	}
	return
}
