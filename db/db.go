package db

import (
	"container/list"
	"database/sql"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/jedib0t/go-pretty/v6/table"
	_ "github.com/mattn/go-sqlite3"
)

// DB file
var DB = "/.get.db"
var MAX_LEN = 80

type Command struct {
	id      int
	name    string
	command string
}

func checkDBPath() bool {
	if _, err := os.Stat(getDB()); err == nil {
		return true
	} else {
		return false
	}
}

func getDB() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		os.Exit(0)
	}
	return homeDir + DB
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
	database, err := sql.Open("sqlite3", getDB())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	initCommandsTable, _ := database.Prepare("CREATE TABLE IF NOT EXISTS commands (id INTEGER PRIMARY KEY, name TEXT, command TEXT)")
	initCommandsTable.Exec()
}

func GetAllCommands() {
	checkDB()
	database, _ := sql.Open("sqlite3", getDB())
	allCommands, _ := database.Query("SELECT * FROM commands")
	var id int
	var name string
	var command string
	arr := list.New()
	for allCommands.Next() {
		allCommands.Scan(&id, &name, &command)
		arr.PushBack(Command{
			id:      id,
			name:    name,
			command: command,
		})
	}

	commandsTable := table.NewWriter()
	commandsTable.SetOutputMirror(os.Stdout)
	commandsTable.SetAllowedRowLength(MAX_LEN)
	commandsTable.SetStyle(table.StyleDouble)
	commandsTable.AppendHeader(table.Row{"#", "Name", "Command"})

	for e := arr.Front(); e != nil; e = e.Next() {
		commandsTable.AppendRow([]interface{}{e.Value.(Command).id, e.Value.(Command).name, e.Value.(Command).command})
	}

	commandsTable.Render()

}

func SetNewCommand(name string, command string) {
	checkDB()
	database, _ := sql.Open("sqlite3", getDB())
	newCommand, _ := database.Prepare("INSERT INTO commands (name, command) VALUES (?,?)")
	newCommand.Exec(name, command)
}

func ShowSpesificCommand(commandIdentifier string) {
	checkDB()
	database, _ := sql.Open("sqlite3", getDB())
	commandToMatch, _ := database.Query("SELECT * FROM commands WHERE id=? OR name=?", commandIdentifier, commandIdentifier)
	var id int
	var name string
	var command string
	arr := list.New()
	for commandToMatch.Next() {
		commandToMatch.Scan(&id, &name, &command)
		arr.PushBack(Command{
			id:      id,
			name:    name,
			command: command,
		})
	}
	commandsTable := table.NewWriter()
	commandsTable.SetOutputMirror(os.Stdout)
	commandsTable.SetAllowedRowLength(MAX_LEN)
	commandsTable.SetStyle(table.StyleDouble)
	commandsTable.AppendHeader(table.Row{"#", "Name", "Command"})
	for e := arr.Front(); e != nil; e = e.Next() {
		commandsTable.AppendRow([]interface{}{e.Value.(Command).id, e.Value.(Command).name, e.Value.(Command).command})
	}
	commandsTable.Render()

}

func DeleteSpesificCommand(commandIdentifier string) {
	checkDB()
	database, _ := sql.Open("sqlite3", getDB())
	commandToDelete, _ := database.Prepare("DELETE FROM commands WHERE id=? OR name=?")
	commandToDelete.Exec(commandIdentifier, commandIdentifier)
}

func CopySpesificCommand(commandIdentifier string) {
	checkDB()
	database, _ := sql.Open("sqlite3", getDB())
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
		err := os.Remove(getDB())
		if err != nil {
			fmt.Println("Failed to reset DB")
		}
	}
	checkDB()
}
