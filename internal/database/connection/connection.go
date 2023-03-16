package connection

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql" // import for side effect
	log "github.com/sirupsen/logrus"
)

// GetConnection returns the db connection instance or error if any
func GetConnection() (*sql.DB, error) {
	url := os.Getenv("DATABASE_URL")
	driver := os.Getenv("DB_DRIVER")

	db, err := sql.Open(driver, url)
	if err != nil {
		log.Error("GetConnection error: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error("GetConnection error: ", err)
		return nil, err
	}

	return db, nil
}
