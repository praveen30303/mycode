package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

func main(){
	file,err := os.Open("zelda.txt")
	if err!=nil{
		log.Fatalf("Failed opening file :%s ",err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtLines[] string

	for scanner.Scan(){
		txtLines = append(txtLines,scanner.Text())
	}
	file.Close()

	for _,linetext := range txtLines {
		fmt.Println(linetext)
	}
}
