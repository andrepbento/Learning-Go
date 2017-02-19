package main

import (
    "fmt"
    "strings"
)

func main() {
    /*      Go strings manipulation libraries
        unicode
        regexp
        strings
    */
    var greeting = "Hello World!"

    fmt.Printf("normal string: %s\n", greeting)
    fmt.Printf("hex bytes: ")
    for i := 0; i < len(greeting); i++ {
        fmt.Printf("%x ", greeting[i])
    }
    fmt.Printf("\n")

    const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
    /*q flag escapes unprintable characters, with + flag it escapses non-ascii characters as well 
     to make output unambigous  */
    fmt.Printf("quoted strings: %+q\n", sampleText)

    fmt.Printf("String Lenght, of variable greeting is: %d\n", len(greeting))

    //  Concatenating Strings
    greetings := []string{"Hello","testeXPTO","world!"}
    fmt.Printf("%s", strings.Join(greetings, ", "))
}