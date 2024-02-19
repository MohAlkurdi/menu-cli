package main

import "fmt"

func main() {

	myBill := newBill("Moh Bill")

	fmt.Println(myBill.format())

}
