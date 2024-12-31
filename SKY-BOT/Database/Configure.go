package Database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // SQLite driver works for DuckDB
	"github.com/stanwar/sky-bot/Models"
)

var DB *gorm.DB

var DB_FILE = "./duckdb_database.db"

func InitializeDB() {

	var err error
	DB, err = gorm.Open("sqlite3", DB_FILE) // Use SQLite driver
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	fmt.Println("Database connected successfully!")

	// creating tables for queued, inprogress and completed tasks(jobs)
	Models.TaskTableName = "queued_jobs"
	DB.AutoMigrate(&Models.Task{})
	Models.TaskTableName = "inprogress_jobs"
	DB.AutoMigrate(&Models.Task{})
	Models.TaskTableName = "completed_jobs"
	DB.AutoMigrate(&Models.Task{})

}
