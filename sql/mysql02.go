package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type City struct{
	Id int
	Name string
	Population int
}
func main(){
	db,err := sql.Open("mysql","student:playstation5@tcp(127.0.0.1:3306)/testdb")
	defer db.Close()

	if err != nil{
		log.Fatal(err)
	}

	res, err2:= db.Query("SELECT * FROM cities")
	defer res.Close()

	if err2!= nil {
		log.Fatal(err2)
	}
	for res.Next(){
		var city City
		err := res.Scan(&city.Id,&city.Name,&city.Population)
		if err!=nil{
			log.Fatal(err)
		}
		fmt.Printf("%v\n",city)

	}
}
