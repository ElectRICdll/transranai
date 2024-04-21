package main

import (
	"os"
	"translator/commands"
)

func main() {
	var command = os.Args[1]
	var params = os.Args[2:]

	commands.CommandSet[command].Execute(params...)
}