package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

type InputParams struct {
	SizeURL        int
	SizeGoroutines int
	API            APIParams
}

type APIParams struct {
	Type     string
	Category string
}

const (
	baseURL           = "https://api.waifu.pics/"
	defaultType       = "sfw"
	defaultCategory   = "waifu"
	defaultURLCount   = 1
	defaultGoroutines = 1
)

func main() {
	var wg sync.WaitGroup
	var input InputParams

	displayInput(&input)

	ch := make(chan string, input.SizeURL)

	tasksPerGoroutine := input.SizeURL / input.SizeGoroutines
	extraTasks := input.SizeURL % input.SizeGoroutines

	for i := 0; i < input.SizeGoroutines; i++ {
		wg.Add(1)
		tasks := tasksPerGoroutine
		if i < extraTasks {
			tasks++
		}
		go fetchURL(ch, &wg, tasks, input.API)
	}

	wg.Wait()
	close(ch)

	displayResults(ch)
}

func displayInput(input *InputParams) {
	fmt.Print("Type (default=sfw): ")
	fmt.Scanln(&input.API.Type)
	if input.API.Type == "" {
		input.API.Type = defaultType
	}

	fmt.Print("Category (default=waifu): ")
	fmt.Scanln(&input.API.Category)
	if input.API.Category == "" {
		input.API.Category = defaultCategory
	}

	fmt.Print("How many URLs (default=1): ")
	fmt.Scanln(&input.SizeURL)
	if input.SizeURL <= 0 {
		input.SizeURL = defaultURLCount
	}

	fmt.Print("How many goroutines (default=1): ")
	fmt.Scanln(&input.SizeGoroutines)
	if input.SizeGoroutines <= 0 {
		input.SizeGoroutines = defaultGoroutines
	}
}

func displayResults(ch chan string) {
	i := 1

	fmt.Println("\n==================================================")
	for url := range ch {
		fmt.Printf("\nURL %d: %s\n", i, url)
		i++
	}
	fmt.Println("==================================================")
}

func fetchURL(ch chan string, wg *sync.WaitGroup, tasks int, apiParams APIParams) {
	defer wg.Done()

	for i := 0; i < tasks; i++ {
		response, err := http.Get(baseURL + apiParams.Type + "/" + apiParams.Category)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			continue
		}
		defer response.Body.Close()

		data, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			continue
		}

		ch <- string(data)
	}
}
