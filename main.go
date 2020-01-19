package main

import "fmt"

func Myout() string {
	fmt.Println("here")
	return "good"
}

func main() {
	fmt.Println("main")
	fmt.Println(Myout())
	fmt.Println("end")
}
