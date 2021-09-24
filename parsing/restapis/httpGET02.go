package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	url := "https://alta3.com"
	method:="GET"

	client:= &http.Client{}

	req,err := http.NewRequest(method,url,nil)

	if err != nil{
		fmt.Println(err)
		return
	}

	res,err := client.Do(req)

	if err!=nil{
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	f,err := os.Create("alta3index.html")


	if err != nil{
		fmt.Println(err)
		return
	}
	defer f.Close()

	_,err = f.ReadFrom(res.Body)
	if err != nil {
		fmt.Println(err)
	}

}
