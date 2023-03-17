package connection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // import for side effect
	log "github.com/sirupsen/logrus"
)

// GetConnection returns the db connection instance or error if any
func GetConnection() (*sql.DB, error) {
	// get connection details from env
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", userName, password, host, port, dbName)

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
