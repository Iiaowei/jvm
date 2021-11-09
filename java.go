package main

import (
	"flag"
	"fmt"
	"jvm/cmd"
)

func main() {
	help := flag.Bool("help", false, "帮助信息")
	version := flag.Bool("version", false, "1.0")
	cp := flag.String("cp", "", "classpath")
	flag.Parse()
	fmt.Println(*cp)

	command := cmd.NewCommand()
	if *help {
		command.PrintHelp()
	}
	if *version {
		command.PrintVersion()
	}
}
