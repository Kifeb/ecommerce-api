package helpers

import "database/sql"

func CommirOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			panic(errRollback)
		}
		panic(err)
	} else {
		tx.Commit()
	}

}
