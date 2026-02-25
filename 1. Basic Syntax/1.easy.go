// -----------------------------------------------------------------------------------------------------------
// QUESTION 1:
// WRITE A MINIMAL GO PROGRAM THAT PRINTS YOUR NAME AND THE TOTAL NUMBER OF ARGUMENTS PASSED VIA THE TERMINAL WHEN EXECUTING THE PROGRAM.
// -----------------------------------------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

func PrintNameAndArguments() {
	fmt.Println("Muhammad Ismail")
	fmt.Println("Total Number of Args = ", len(os.Args))
}
