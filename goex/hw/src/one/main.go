package main

import (

	"fmt"

)

func main (){

	c:= make(chan int)
	var nums int
	fmt.Println("how many iterations would you like")
	fmt.Scanln(&nums)
	go sender(c, nums)
	go receiver(c)

	var input string

	fmt.Scanln(&input)


}

