package main

import "fmt"

func main(){
	//artimatik operator

	//penjumlahan
	a := 600
	b := 300

	fmt.Println("Penjumlahan " , a+b);
	fmt.Println("Pengurangan " , a/b);
	fmt.Println("Perkalian " , a*b);
	fmt.Println("Pembagian " , a/b);
	fmt.Println("Modulus " , a%b);

	//operator relasional => bool (true, false)
	i := 10
	j := 5

	fmt.Println(i != j) //TRUE
	fmt.Println(i == j) //FALSE
	fmt.Println(i > j) //TRUE
	fmt.Println(i < j) //FALSE
}