package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type MSSQLConfig struct {
	User     string
	Password string
	Address  string // host:port
	DBName   string
}

func NewMSSQLStorage(cfg MSSQLConfig) (*sql.DB, error) {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s",
		cfg.User, cfg.Password, cfg.Address, cfg.DBName)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Println("❌ Failed to open DB connection:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("❌ Failed to ping DB:", err)
		return nil, err
	}

	log.Println("✅ Successfully connected to SQL Server!")
	return db, nil
}
