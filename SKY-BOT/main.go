package main

import (
	"fmt"

	"github.com/stanwar/sky-bot/Database"
	"github.com/stanwar/sky-bot/Routes"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	Database.InitializeDB()
	defer Database.DB.Close()

	fmt.Println("Server started!!!")
	r := Routes.SetupRouter()
	r.Run()

}
