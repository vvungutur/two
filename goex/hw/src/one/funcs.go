package main



import (

	"fmt"
	"time"

)


func sender(c chan int, n int){

	for i := 0; i <= n; i ++{

		c <- i
		time.Sleep(time.Second * 1)
	}


}

func receiver(c chan int){

	for {
		msg:= <- c
		fmt.Println(msg)

	}

}
