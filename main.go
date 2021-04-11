package main

import (
	"devotedbapp/cli"
	"devotedbapp/database"

	_ "github.com/proullon/ramsql/driver"
)

func main() {
	database.SetupDatabase()

	cli.ProcessInterective()
}
