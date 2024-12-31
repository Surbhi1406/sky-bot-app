package Database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/stanwar/sky-bot/Models"
)

func InsertTask(task *Models.Task) error {
	// insert into db
	if err := DB.Create(&task).Error; err != nil {
		fmt.Println("Error inserting task:", err)
	} else {
		fmt.Println("Task inserted successfully:", task)
	}
	return nil
}

// Insert a record into a table
func InsertTaskIntoTable(db *gorm.DB, tableName string, task Models.Task) error {
	return db.Table(tableName).Create(&task).Error
}

func FetchTaskFromTable(db *gorm.DB, tableName string, id string) (Models.Task, error) {
	var task Models.Task
	err := db.Table(tableName).First(&task, "job_id = ?", id).Error
	return task, err
}

// FetchAllTasks fetches all entries from the 'task' table.
func FetchAllTasks(task *[]Models.Task) error {
	// var tasks []Models.Task

	// Query the 'task' table
	if err := DB.Find(&task).Error; err != nil {
		fmt.Println("Error fetching tasks:", err)
		return err
	}

	fmt.Println("Fetching Data from db")
	fmt.Println(task)
	return nil
}

// Delete a record from a table
func DeleteTaskFromTable(db *gorm.DB, tableName string, id string) error {
	return db.Table(tableName).Delete(&Models.Task{}, "job_id = ?", id).Error
}

// Move a record from one table to another
func MoveTaskBetweenTables(db *gorm.DB, fromTable string, toTable string, id string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		// Fetch task from source table
		task, err := FetchTaskFromTable(tx, fromTable, id)
		if err != nil {
			return err
		}

		// Insert task into target table
		if err := InsertTaskIntoTable(tx, toTable, task); err != nil {
			return err
		}

		// Delete task from source table
		if err := DeleteTaskFromTable(tx, fromTable, id); err != nil {
			return err
		}

		return nil
	})
}

func GetFirstQueuedJob(db *gorm.DB, tableName string) (Models.Task, error) {
	var task Models.Task
	// Fetch the task with the earliest created_at timestamp
	err := db.Table(tableName).Order("created_at ASC").First(&task).Error
	return task, err
}
