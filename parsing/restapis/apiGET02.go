package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"time"
)

type SpaceX []struct {
	CapsuleSerial      string    `json:"capsule_serial"`
	CapsuleID          string    `json:"capsule_id"`
	Status             string    `json:"status"`
	OriginalLaunch     time.Time `json:"original_launch"`
	OriginalLaunchUnix int       `json:"original_launch_unix"`
	Missions           []struct {
		Name   string `json:"name"`
		Flight int    `json:"flight"`
	} `json:"missions"`
	Landings   int    `json:"landings"`
	Type       string `json:"type"`
	Details    string `json:"details"`
	ReuseCount int    `json:"reuse_count"`
}

func main(){
	url := "https://api.spacexdata.com/v3/capsules"
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

	var record SpaceX

	if err:= json.NewDecoder(res.Body).Decode(&record); err!=nil {
		log.Println(err)
	}

	for launchNo, launchData := range record {
		fmt.Println("Capsule Record     =\n", launchData)
        fmt.Println("Record Number      =", launchNo)
        fmt.Println("Capsule Serial     =", launchData.CapsuleSerial)
	}
}
