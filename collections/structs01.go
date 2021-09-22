package main

import "fmt"

type Coord struct{
	x int
	y int
}

func main(){
	fmt.Println(Coord{1,2})
	z:= Coord{42,100}
	z.y=180
	fmt.Println(z)
	fmt.Println("The X coordinate is:", z.x)
    fmt.Println("The Y Coordinate is:", z.y)

}
