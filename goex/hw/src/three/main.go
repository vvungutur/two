package main


import (

	"fmt"
	"io/ioutil"
	"net/http"

)

func getPage(url string) (int, error) {
	
	resp, err := http.Get(url)

	if err != nil{
		return 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		return 0, err

	}

	return len(body), nil
}

func worker (urlChan chan string, c chan string, id int){
	
	url := <- urlChan
	
	length, err := getPage(url)
	if err == nil{
		
		c <- fmt.Sprintf("The length of %s is %d (process no. %d)",  url, length, id)

	} else {
		c <- fmt.Sprintf("is fucked")
	}

}


func main() {
  

  urls := []string {"https://google.com", "https://www.yahoo.com", "https://www.zazu.com", "https://www.github.com", "https://www.facebook.com"}

	d := make(chan string)
	c := make(chan string)

	for i := 0; i < len(urls); i ++{

	go worker(d, c, i)


	}

	for _, url := range urls {	
	
	d <- url


	}

	for i:= 0; i < len(urls); i ++ {
		fmt.Printf("%s\n", <- c)
	}
	
}
