package main

import (
	"fmt"
	"jvm/src/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

//
//	startJVM
//	@Description: start the jvm
//	@param cmd *Cmd
//
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classDate, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Counld not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classDate)
}
