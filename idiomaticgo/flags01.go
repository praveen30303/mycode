package main

import (
	"flag"
	"fmt"
)

func main(){
	num := flag.Int("n",5,"How angry is the Hulk? (# of iterations)")
	flag.Parse()
	n:=*num

	for i:=0;i<n;i++{
		fmt.Println("Hulk Smash!!")
	}
}
