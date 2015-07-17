package main 


import (

	"fmt"
	"time"	

)


func sender(c chan int){


	for i:= 0; i <= 3; i ++ {

		c <- i
		time.Sleep(time.Second * 1)


	}


}

func receiver(c chan int){


	for  {

		msg :=  <- c
		fmt.Println(msg)


	}


}
func main(){

	c := make(chan int)

	go sender(c)
	go receiver(c)

	outside()
	var input string

	fmt.Scanln(&input)

}
