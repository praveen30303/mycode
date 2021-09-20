package main

import (
	"fmt"
	"strconv"
)

func main(){
	s:="33"
	i,err := strconv.Atoi(s)
	if err!= nil {
		panic (err)
	}
	fmt.Println("string:",s)
	fmt.Println("Integer:",i)

	st:= strconv.Itoa(42)
	fmt.Printf("st is now of type , %T",st)
	fmt.Println("String",st)
}
