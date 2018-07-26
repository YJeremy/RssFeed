package search

import (
	"log"
)

// Result contains the result of a search.
type Result struct {
	Field   string
	Content string
	Date 	string
}
//类型 ：title or description
//内容 ： 

//定义Matcher接口类型，接口需要满足拥有方法Search()
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently.
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

// Display writes results to the console window as they
// are received by the individual goroutines.
func Display(results chan *Result)(string) {
	// The channel blocks until a result is written to the channel.
	// Once the channel is closed the for loop terminates.
	var txt string 
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content,result.Date)
		txt = txt + result.Field +"\r\n" + "\r\n" + result.Content +  "\r\n" + "\r\n" + result.Date +  "\r\n" + "\r\n" 
	}

	return txt
}
