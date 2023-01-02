package helpers

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx) error {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			// panic(errRollback)
			return errRollback
		}
		// panic(err)
		fmt.Println(err)
	} else {
		tx.Commit()
	}
	return nil
}
