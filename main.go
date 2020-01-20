package main

import "fmt"

func Myout() string {
	fmt.Println("there")
	return "good"
}

func main() {
	fmt.Println("main")
	fmt.Println(Myout())
	fmt.Println("end")
}
