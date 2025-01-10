package main

import (
	"go-quiz/cmd/api"
	"go-quiz/cmd/cli"
)

func main() {
	api.Start()
	cli.Start()
}
