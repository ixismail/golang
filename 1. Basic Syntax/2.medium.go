// -----------------------------------------------------------------------------------------------------------
// QUESTION 2:
// EXPLAIN WHY THE FOLLOWING CODE SNIPPET FAILS TO COMPILE, AND REWRITE IT SO THAT IT COMPILES AND RUNS CORRECTLY
// -----------------------------------------------------------------------------------------------------------

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

package main
