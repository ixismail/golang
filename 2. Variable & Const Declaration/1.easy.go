// -----------------------------------------------------------------------------------------------------------
// QUESTION 1:
// IDENTIFY WHICH OF THE FOLLOWING DECLARATIONS ARE VALID AT THE PACKAGE LEVEL (OUTSIDE OF ANY FUNCTION), AND WHICH CAUSE COMPILE ERRORS:
// 		var dbHost string = "localhost"
// 		dbPort := 5432
// 		const maxConnections = 100
// 		var isReady bool
// -----------------------------------------------------------------------------------------------------------

package main

var dbHost string = "localhost" // CORRECT

// dbPort := 5432 INCORRECT WE CANT USE := FOR GLOBAL VARIABLES

const maxConnections = 100 // CORRECT

var isReady bool // CORRECT
