package main


import (

	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

func main() {

	urls := []string {"https://google.com", "https://www.yahoo.com", "https://www.zazu.com", "https://www.github.com", "https://www.gangster.com"}

	for _, url :=  range urls{

		pageLength, err := getPage(url)

		if err != nil{
			os.Exit(1)
		}

		fmt.Printf("%s is length %d\n", url, pageLength)
	}
}
