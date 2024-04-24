package main

import (
	"os"
	"translator/commands"
)

func main() {
	commands.CommandSelector(os.Args[1:]...).Execute()
}