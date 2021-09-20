package main

import (
	"fmt"
	"math/cmplx"
)

var (
	kirk bool = false
	spock uint64 = 1<<64-1
	mccoy complex128 = cmplx.Sqrt(-5+12i)
)
func main(){
	fmt.Printf("Type: %T Value: %v\n", kirk, kirk)
	fmt.Printf("Type: %T Value: %v\n", spock, spock)
	fmt.Printf("Type: %T Value: %v\n", mccoy, mccoy)
}
