package main

import "fmt"

func main() {
    array := [5]int{1, 2, 3, 4, 5}
    slice:= array[:]
    fmt.Println("Modifying Slice")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)

    array[1] = 111
    slice[3]=333
    fmt.Println("Modifying Slice")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)
}	
