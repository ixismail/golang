// -----------------------------------------------------------------------------------------------------------
// QUESTION 2:
// LOOK AT THE FOLLOWING FUNCTION THAT ATTEMPTS TO UPDATE SERVER
// CONNECTION VARIABLES USING THE SHORT DECLARATION SYNTAX (:=):
/*
	package main

	import "fmt"

	func main() {
		host := "localhost"

		// Attempt A
		host, port := "127.0.0.1", 8080

		// Attempt B
		host, port := "0.0.0.0", 9000

		fmt.Println(host, port)
	}
*/
// 1. DOES ATTEMPT A COMPILE? WHY OR WHY NOT?
// 2. DOES ATTEMPT B COMPILE? WHY OR WHY NOT?
// 3. HOW SHOULD ATTEMPT B BE WRITTEN SO THE CODE COMPILES AND UPDATES BOTH VARIABLES?
// -----------------------------------------------------------------------------------------------------------

// ANS 1: YES, BEACUSE THE port VARIABLE IS THE NEW ONE SO port WILL BE CREATED AND host WILL BE REUSED
// ANS 2: NO, BECAUSE NOW THERE NO NEW VARIABLE ON THE LEFT SIDE OF THE :=s
// ANS 3: CORRECT ATTEMPT B

package main

import "fmt"

func Medium() {
	host := "localhost"

	host, port := "127.0.0.1", 8080

	// Attempt B (USE = INSTEAD OF :=)
	host, port = "0.0.0.0", 9000

	fmt.Println(host, port)
}
