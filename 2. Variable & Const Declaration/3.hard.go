// -----------------------------------------------------------------------------------------------------------
// QUESTION 3:
// SPOT THE BUG IN THE FOLLOWING BACKEND CONFIGURATION LOADING LOGIC,
// EXPLAIN WHAT VALUE dbTimeout WILL HAVE AT THE END, AND WRITE THE CORRECTED VERSION:
/*
	package main

	import (
		"fmt"
		"strconv"
	)

	func parseTimeout(s string) (int, error) {
		return strconv.Atoi(s)
	}

	func main() {
		var dbTimeout int = 5 // default 5 seconds
		rawInput := "15"

		if rawInput != "" {
			dbTimeout, err := parseTimeout(rawInput)
			if err != nil {
				fmt.Println("Error parsing timeout:", err)
				return
			}
			fmt.Println("Parsed timeout inside if:", dbTimeout)
		}

		fmt.Println("Final configured timeout:", dbTimeout)
	}
*/
// -----------------------------------------------------------------------------------------------------------

// THE LOGIC BUG HERE IS THAT THE dbTimeout IS DECLARED OUTSIDE THE SCOPE OF IF
// BUT THE := SYNTAX WILL CREATE A NEW LOCAL dbTimeout VARIABLE INSTEAD OF z
// ASSIGNING TO THE UPPER ONE SO THE OUTPUT FOR THIS WILL BE:
// OUTPUT:
// 		Parsed timeout inside if: 15
// 		Final configured timeout: 5

// =============== CORRECT CODE ===============

package main

import (
	"fmt"
	"strconv"
)

func parseTimeout(s string) (int, error) {
	return strconv.Atoi(s)
}

func Hard() {
	var dbTimeout int = 5 // default 5 seconds
	rawInput := "15"

	if rawInput != "" {
		var err error
		dbTimeout, err = parseTimeout(rawInput)
		if err != nil {
			fmt.Println("Error parsing timeout:", err)
			return
		}
		fmt.Println("Parsed timeout inside if:", dbTimeout)
	}

	fmt.Println("Final configured timeout:", dbTimeout)
}
