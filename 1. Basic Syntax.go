package main

import (
	"fmt"
	"os"
)

// QUESTION 1: WRITE A MINIMAL GO PROGRAM THAT PRINTS YOUR NAME AND THE TOTAL NUMBER OF ARGUMENTS PASSED VIA THE TERMINAL WHEN EXECUTING THE PROGRAM.

func printNameAndArguments() {
	fmt.Println("Muhammad Ismail")
	fmt.Println("Total Number of Args = ", len(os.Args))
}

// QUESTION 2: EXPLAIN WHY THE FOLLOWING CODE SNIPPET FAILS TO COMPILE, AND REWRITE IT SO THAT IT COMPILES AND RUNS CORRECTLY
/*
	func main()
	{
    	currentTime := time.Now()
    	status := "running"
    	fmt.Println("Application started")
	}
*/

// ANS: THE CODE FAILS FOR TWO REASONS
// 1. FUNCTION DECLARATIONS IS WRONG (OPENING BRACKET MUST BE ON SAME LINE AS FUNCTION NAME)
// 		CORRECT = func main() {
// 2. UNUSED LOCAL VARIABLES (currentTime, status)

// QUESTION 3: IN BACKEND APPLICATIONS, SERVICES FREQUENTLY ACCEPT COMMAND-LINE FLAGS OR ARGUMENTS AT STARTUP. WRITE A PROGRAM WITH PACKAGE MAIN THAT ITERATES OVER OS.ARGS (SKIPPING THE EXECUTABLE PATH AT INDEX 0) AND PRINTS EACH ARGUMENT ALONGSIDE ITS 1-BASED INDEX (E.G., ARG 1: FOO, ARG 2: BAR). IF NO ARGUMENTS ARE PROVIDED, IT SHOULD PRINT "NO ARGUMENTS PROVIDED" AND TERMINATE WITH AN EXIT STATUS CODE OF 1.

func printArgsWithIndex() {
	if len(os.Args) <= 1 {
		fmt.Println("No Arguments provided!")
		os.Exit(1)
	}

	for i, v := range os.Args[1:] {
		fmt.Printf("ARG %d: %s  ", i, v)
	}
}

func main() {
	printNameAndArguments()
	printArgsWithIndex()
}
