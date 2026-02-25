// -----------------------------------------------------------------------------------------------------------
// QUESTION 3:
// IN BACKEND APPLICATIONS, SERVICES FREQUENTLY ACCEPT COMMAND-LINE FLAGS OR ARGUMENTS AT STARTUP. WRITE A PROGRAM WITH PACKAGE MAIN THAT ITERATES OVER OS.ARGS (SKIPPING THE EXECUTABLE PATH AT INDEX 0) AND PRINTS EACH ARGUMENT ALONGSIDE ITS 1-BASED INDEX (E.G., ARG 1: FOO, ARG 2: BAR). IF NO ARGUMENTS ARE PROVIDED, IT SHOULD PRINT "NO ARGUMENTS PROVIDED" AND TERMINATE WITH AN EXIT STATUS CODE OF 1.
// -----------------------------------------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

func PrintArgsWithIndex() {
	if len(os.Args) <= 1 {
		fmt.Println("No Arguments provided!")
		os.Exit(1)
	}

	for i, v := range os.Args[1:] {
		fmt.Printf("ARG %d: %s  ", i, v)
	}
}
