package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SeoData struct{}
type Parser interface{}

// setting user agents
var userAgents = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	fmt.Println(userAgents[randNum])
	return userAgents[randNum]
}

// site map url extractor
func extractSiteMapUrls() {
	makeRequest()
}

// request maker
func makeRequest() {}

// scrape url
func scrapeUrl() {}

// scrape Page
func scrapePage() {}

// crawl page
func crawlPage() {}

// get seo data
func getSeoData() {}

// scrape sitemap
func scrapeSitemap() {
	scrapeUrl()
	extractSiteMapUrls()
}
func main() {

}
