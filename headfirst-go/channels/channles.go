package channels

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}

func even(channel chan int) {
	channel <- 2
	channel <- 4
}

func Test() {
	channelA := make(chan int)
	channelB := make(chan int)
	go odd(channelA)
	go even(channelB)
	fmt.Println(<-channelA)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)
	fmt.Println(<-channelB)
}

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func FetchHttpResponseSizes(urls []string) {
	pages := make(chan Page)

	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}
