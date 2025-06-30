package main

type WebsiteChecker func(string) bool

type checkResult struct {
	url       string
	isSuccess bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	resultChannel := make(chan checkResult)
	for _, url := range urls {
		go func() {
			isSuccess := wc(url)
			resultChannel <- checkResult{url: url, isSuccess: isSuccess}
		}()
	}

	results := make(map[string]bool)
	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.isSuccess
	}
	return results
}
