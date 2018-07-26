package search

import (
	"log"
	"sync"
	"os"
	"fmt"
)

// A map of registered matchers for searching.
// 建立一个匹配器存储的容器，命名和值，如rss字符串为键名，rssMatcher类型的变量为键值
var matchers = make(map[string]Matcher)

// Run performs the search logic.
// 实现主要逻辑功能
func Run(searchTerm string){
	// Retrieve the list of feeds to search through.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while
	// they process the individual feeds.
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		// Retrieve a matcher for the search.
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed.
	txt := Display(results)
	wiriteFile(txt)
	
}

// Register is called to register a matcher for use by the program.
// 注册器，传入注册的匹配器类型 matcher，以及该类型的名字feedType
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	//存入 匹配器的映射，建立键值对；该匹配器的命名和值
	matchers[feedType] = matcher
}

func wiriteFile(body string){

	var data[]byte = []byte(body)
	fout, err := os.Create("RSS.md")
	if err !=nil{
		fmt.Println(err)
		return	
	}
	defer fout.Close()
	fout.Write(data)
	return
}