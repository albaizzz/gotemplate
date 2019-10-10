package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

//SQLDb base contract
type SQLDb interface {
	BeginMasterx() *sqlx.Tx
}

//rebindIn rebinds the query with params using IN
//Translate query using IN
//Example :
//cartIDs := []int64{2222, 3333, 44444}
//userID := 1066513
//Select * from ws_shopping_cart where cart_id IN ($1) and user_id = ($2).
// Will Be Process like this :
// param = []int64{2222, 3333, 44444, 1066513}
// Select * from ws_shopping_cart where cart_id IN (?, ?, ?) and user_id = (?).
func rebindIn(queryS string, param ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(queryS, param...)
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func execute(ctx context.Context, dbsqlx *sqlx.DB, tx *sqlx.Tx, query string, params ...interface{}) (int64, error) {
	query, args, err := rebindIn(query, params...)
	if err != nil {
		return 0, err
	}
	stmt, err := tx.PrepareContext(ctx, dbsqlx.Rebind(query))
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return 0, err
	}
	success, err := result.RowsAffected()
	return success, err
}
