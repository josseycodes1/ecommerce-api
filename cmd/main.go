package main

import (
	"database/sql"
	"log"

	"github.com/josseycodes1/ecom-api/cmd/api"
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

	initStorage(db)

	// Start your API server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal("❌ Error running API server:", err)
	}
}

func initStorage(db *sql.DB) {
	// Ping the DB to ensure a successful connection
	err := db.Ping()
	if err != nil {
		log.Fatal("❌ Cannot ping DB:", err)
	}

	log.Println("✅ DB: Successfully connected!")
}
