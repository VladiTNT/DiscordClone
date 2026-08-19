package db

import (
	"database/sql"

	"github.com/VladiTNT/DiscordClone/internal/db/models"
	_ "modernc.org/sqlite"
)

func New(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", databaseUrl)
	if err != nil {
		return nil, err
	}

	err = MigrateDB(db,
		models.UserTable,
		models.PhotoTable,
	)

	return db, err
}

func MigrateDB(db *sql.DB, tables ...string) error {
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return err
		}
	}
	return nil
}
