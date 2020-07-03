package main
// defines which code bundle this file belongs to

import (
    "fmt" // to formatting and printing functions
    "strings" // to manipulate strings
)

func main() {

    fmt.Println("Please enter your name.")

    var name string // create new variable called name of type string

    fmt.Scanln(&name) 
    /* 
     * capture user's input and convert to a string
     * assign them into the var called name
     */

    name = strings.TrimSpace(name) // removes any space and \n characters

    fmt.Printf("Hi, %s! I'm Go!", name)
    /*
     * takes a string, injects var name into the string using special 
     * printing verbs %s
     */
}
