package main

import (
	"flag"
	"fmt"
)

type AddCommand struct {
}

func (c *AddCommand) Synopsis() string {
	return "Add todo task to list"
}

func (c *AddCommand) Help() string {
	return "Usage: todo add [option]"
}

func (c *AddCommand) Run(args []string) int {
	// TODO

	const defaultPort = 3000
	var port int
	f := flag.NewFlagSet("add", flag.ExitOnError)
	f.IntVar(&port, "port", defaultPort, "port to use")
	f.IntVar(&port, "p", defaultPort, "port to use (short)")
	f.Parse(args)
	fmt.Println(port)
	return 0

}
