package main

import (
    "fmt"
)

func main() {
    /*      ARRAY DECLARATION
        var variable_name [SIZE] variable_type
    */

    var balance [10] float32
    fmt.Println("\nvar balance")
    for i, element := range balance {
        fmt.Printf("%d : %.4f\n", i, element)
    }
    // The size of the array can not be larger than the number of elements that we declare
    var balance1 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    fmt.Println("\nvar balance1")
    for i := range balance1 {
        fmt.Printf("%d : %.4f\n",i , balance1[i])
    }
    // If we omit the size of the array, an array just big enough to hold the initialization is created
    //var balance2 = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    //fmt.Printf("%.2f", balance2[5])  --> nil

    var n [10]int // n is an array of 10 integers
    for i := 0; i < 10; i++ {
        n[i] = i + 100 // set element at location i to i + 100
    }

    fmt.Println("\nvar n")
    for i, element := range n {
        fmt.Printf("Element[%d] = %d\n", i, element)
    }

    //      Multi-Dimensional Arrays
    //  var variable_name [SIZE10][SIZE2]...[SIZEN] variable_type
    //      Example:
    //  var threedim [5][10][4]int
    //  var arrayName [x][y] variable_type
    //      Initializing two-Dimensional Arrays:
    var a = [3][4]int {
        {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
        {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
        {8, 9, 10, 11} }   /*  initializers for row indexed by 2 */
    fmt.Println("\nMulti-Dimensional array - var a [3][4]int")
    for i := 0; i < 3; i++ {
        for j := 0; j < 4; j++ {
            fmt.Printf("Element[%d][%d] : %d\n", i, j, a[i][j])
        }
    }

    /*      Passing arrays to functions
        Way-1
        func myFunction(param [10]int) {
            ...
        }
        Way-2
        func myFunction(param []int) {
            ...
        }
    */
    var balance2 = []int{1000, 2, 3, 17, 50}
    fmt.Printf("\nAverage value, of var balance2 is %.2f", getAverage(balance2, len(balance2)))
}

func getAverage(arr []int, size int) float32 {
    var sum int
    for i := 0; i < size; i++ {
        sum += arr[i]
    }
    return float32(sum / size)
}