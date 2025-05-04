package main

import (
	"github.com/andreixhz/sshm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}