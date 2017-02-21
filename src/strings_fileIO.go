package main

import (
    "fmt"
    "strings"
    "sort"
    "os"
    "log"
    "io/ioutil"
)

func main() {
    
    sampString := "Hello World"

    fmt.Println(strings.Contains(sampString, "lo"))
    
    fmt.Println(strings.Index(sampString, "lo"))
    
    fmt.Println(strings.Count(sampString, "l"))

    fmt.Println(strings.Replace(sampString, "l", "x", 3))

    cvsString := "1,2,3,4,5,6"

    fmt.Println(strings.Split(cvsString, ","))

    listOfLetters := []string{"c", "a", "b"}

    sort.Strings(listOfLetters)

    fmt.Println("Letters: ", listOfLetters)

    listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")

    fmt.Println(listOfNums)

    file, err := os.Create("samp.txt")

    if err != nil {
        log.Fatal(err)
    }

    file.WriteString("This is some random text")

    file.Close()

    stream, err := ioutil.ReadFile("samp.txt")

    if err != nil {
        log.Fatal(err)
    }

    readString := string(stream)

    fmt.Println(readString)
}