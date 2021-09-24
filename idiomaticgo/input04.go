package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input,err := reader.ReadString('\n')
	if err!= nil{
		fmt.Println("An error occurred during processing \n")
		return
	}
	input = strings.TrimSuffix(input,"\n")
	fmt.Println(input)
}
