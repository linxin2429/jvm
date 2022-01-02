package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd
// @Description: struct for cmd parse
//
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	// classpath
	cpOption string
	// path to jre
	XjreOption string
	// java class
	class string
	args  []string
}

//
//	parseCmd
//	@Description: cmd parse
//	@return *Cmd
//
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

//
//	printUsage
//	@Description: print the usage of jvm
//
func printUsage() {
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}
