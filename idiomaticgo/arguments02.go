package main

import (
	"fmt"
	"os"
)

func main(){
	argsWithoutProg := os.Args[1:]
		
	for i,a:= range argsWithoutProg {
		fmt.Printf(" arg at pos %d is %s\n",i+1,a)
	}
}
