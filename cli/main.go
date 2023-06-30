package main

import (
	"envdost/cmd"
	_ "envdost/cmd/delete"
)

func main() {
	cmd.Execute()
}