package main 
// defines which code bundle this file belongs to

import "fmt" 
/* 
 * tells the Go compiler which other packages to use in this file
 * "fmt" package provides formatting and printing functions
 */

func main() {
	fmt.Println("Hello, World!")
	// tells the computer to print some text (strings) to the screen
}
/*
 * Go applications require a "main" package and exactly one main() function
 * the main() function takes no arguments and returns no values
 * it tells the Go compiler to compile that package as an executable
 */
