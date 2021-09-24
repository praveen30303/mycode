package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	f,err := os.Open("vendorData.csv")

	if err!=nil{
		log.Fatal(err)
	}

	r:= csv.NewReader(f)

	for {
		fmt.Println("reading line")
		record,err := r.Read()

		if err == io.EOF {
			break
		}

		if err!= nil{
			log.Fatal(err)
		}
		for value:=range record {
			fmt.Printf("%s\n",record[value])
		}
	}


}
