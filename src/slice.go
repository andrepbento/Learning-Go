package main

import (
    "fmt"
)

func main() {
    var numbers = make([]int,3,5) // create a slice of length 3 and capacity 5

    printSlice(numbers)

    /*  If the slice is not inicialized, the len and
        cap are 0 and the default value is nil
    */
    var numbers1 []int
    printSlice(numbers1)
    if(numbers1 == nil) {
        fmt.Printf("slice is nil\n")
    }

    // create a slice
    numbers2 := []int{0,1,2,3,4,5,6,7,8}
    printSlice(numbers2)

    /* print the original slice */
    fmt.Println("numbers2 ==", numbers2)
   
    /* print the sub slice starting from index 1(included) to index 4(excluded) */
    fmt.Println("numbers2[1:4] ==", numbers2[1:4])
   
    /* missing lower bound implies 0 */
    fmt.Println("numbers2[:3] ==", numbers2[:3])
   
    /* missing upper bound implies len(s) */
    fmt.Println("numbers2[4:] ==", numbers2[4:])
   
    numbers3 := make([]int,0,5)
    printSlice(numbers3)
   
    /* print the sub slice starting from index 0(included) to index 2(excluded) */
    numbers4 := numbers2[:2]
    printSlice(numbers4)
   
    /* print the sub slice starting from index 2(included) to index 5(excluded) */
    numbers5 := numbers2[2:5]
    printSlice(numbers5)

    //      append() and copy() functions
    var numbers6 []int
    printSlice(numbers6)

    /* append allows nil slice */
    numbers6 = append(numbers6, 0)
    printSlice(numbers6)

    /* add one element to slice */
    numbers6 = append(numbers6, 1)
    printSlice(numbers6)

    /* add more than one element at a time */
    numbers6 = append(numbers6, 2,3,4)
    printSlice(numbers6)

    /* crate a slice numbers7 with double the capacity of earlier slice */
    numbers7 := make([]int, len(numbers6), (cap(numbers6))*2)

    /* copy content of numbers to numbers7 */
    copy(numbers7, numbers6)
    printSlice(numbers7)

}

func printSlice(x []int) {
    fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
