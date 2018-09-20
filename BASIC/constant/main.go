package main

import "fmt"

const x string = "Berlind"

func main(){
	fmt.Println(x)

	const z int = 20
	fmt.Println(z)

	c := 200 / z
	fmt.Println(c)
}