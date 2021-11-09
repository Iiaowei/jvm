package cmd

import "fmt"

type command struct {
	help string
	cp string
	version string
}

func NewCommand() *command {
	return &command{
		help: "帮助信息",
		version: "Java version 0.1",
	}
}

func (c *command) PrintVersion() {
	fmt.Println(c.version)
}

func (c *command) PrintHelp() {
	fmt.Println(c.help)
}