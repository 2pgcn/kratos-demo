package main

import "fmt"

type a struct {
	aa *string
	bb string
}

func main() {
	a1 := &a{}
	fmt.Println(a1.bb == "")
	fmt.Println(a1.aa == nil)
	b := ""
	a1.aa = &b
	fmt.Println(a1.aa == "")

}
