package main

import (
	"github.com/Elianamos29/go-task-manager-cli/cmd"
	"github.com/Elianamos29/go-task-manager-cli/db"
)

func main() {
	db.InitDB("tasks.db")

	cmd.HandleCommands()
}