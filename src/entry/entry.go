package main

import _ "configuration"
import "flag"
import "fmt"
import "framework"
import "os"

func main() {
	canRunAll, exampleId := handleParameters()
	fmt.Println("Let's Go!")
	if canRunAll {
		framework.RunAll()
	} else {
		framework.RunSingle(exampleId)
	}
}

func handleParameters() (bool, string) {
	runAll := flag.Bool("a", false, "Run all examples")
	runSingle := flag.String("r", "", "Run a single example by specifying the example id")
	parameters := os.Args[1:]
	flag.Parse()
	if len(parameters) < 1 {
		fmt.Println("Usage: /path/to/entry [-h] [-a] [-r exampleId]")
		flag.PrintDefaults()
		os.Exit(-1)
	}

	if *runAll {
		return true, ""
	} else {
		return false, *runSingle
	}
}
