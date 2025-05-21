package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/josseycodes1/ecom-api/config"
	"github.com/josseycodes1/ecom-api/db"
)

func main() {
	// Create a new MSSQL storage with the config values
	mssqlConfig := db.MSSQLConfig{
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		Address:  config.Envs.DBAddress, // Use 'localhost:1433' or the appropriate address
		DBName:   config.Envs.DBName,
	}

	// Initialize DB connection
	db, err := db.NewMSSQLStorage(mssqlConfig)
	if err != nil {
		log.Fatal("❌ Error connecting to DB:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

}
