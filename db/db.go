package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	_ "github.com/mattn/go-sqlite3"
)

// DB file
var DB = "db/get.db"

type Command struct {
	id      int
	name    string
	command string
}

func checkDBPath() bool {
	if _, err := os.Stat(DB); err == nil {
		return true
	} else {
		return false
	}
}

func checkDB() {
	// Check if DB exsist
	if !checkDBPath() {
		// Creating DB
		InitDB()
		fmt.Println("New DB created")
		os.Exit(0)
	}
}

func InitDB() {
	database, err := sql.Open("sqlite3", DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	initCommandsTable, _ := database.Prepare("CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, name TEXT, command TEXT)")
	initCommandsTable.Exec()
}

func GetAllCommands() {
	checkDB()
	database, _ := sql.Open("sqlite3", DB)
	allCommands, _ := database.Query("SELECT * FROM commands")
	var id int
	var name string
	var command string
	for allCommands.Next() {
		allCommands.Scan(&id, &name, &command)
		fmt.Println(id, name, command)
	}
}

func SetNewCommand(name string, command string) {
	checkDB()
	database, _ := sql.Open("sqlite3", DB)
	newCommand, _ := database.Prepare("INSERT INTO commands (name, command) VALUES (?,?)")
	newCommand.Exec(name, command)
}

func ShowSpesificCommand(commandIdentifier string) {
	checkDB()
	database, _ := sql.Open("sqlite3", DB)
	commandToMatch, _ := database.Query("SELECT * FROM commands WHERE id=? OR name=?", commandIdentifier, commandIdentifier)
	var id int
	var name string
	var command string
	for commandToMatch.Next() {
		commandToMatch.Scan(&id, &name, &command)
		fmt.Println(id, name, command)
	}

}

func DeleteSpesificCommand(commandIdentifier string) {
	checkDB()
	database, _ := sql.Open("sqlite3", DB)
	commandToDelete, _ := database.Prepare("DELETE FROM commands WHERE id=? OR name=?")
	commandToDelete.Exec(commandIdentifier, commandIdentifier)
}

func CopySpesificCommand(commandIdentifier string) {
	checkDB()
	database, _ := sql.Open("sqlite3", DB)
	commandToMatch, _ := database.Query("SELECT * FROM commands WHERE id=?", commandIdentifier)
	var id int
	var name string
	var command string
	for commandToMatch.Next() {
		commandToMatch.Scan(&id, &name, &command)
		fmt.Println(command, "- copied to clipboard!")
		break
	}

	// Copy command to clipboard
	clipboard.WriteAll(command)

}

func ResetDB() {
	if checkDBPath() {
		err := os.Remove(DB)
		if err != nil {
			fmt.Println("Failed to reset DB")
		}
	}
	checkDB()
}
