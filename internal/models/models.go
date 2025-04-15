package models

import "database/sql"

func Init(db *sql.DB) {
	Examples.db = db
}
