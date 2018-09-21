package main

import "fmt"

func main(){
	x := 100

	switch x {
	case 60:
		fmt.Println("Nilai mu C")
	case 80 :
		fmt.Println("Nilai mu B")
	case 100 :
		fmt.Println("Nilai mu A")
	default:
		fmt.Println("Nilai mu Gk Jelas")
	}
}