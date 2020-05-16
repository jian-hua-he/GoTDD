package concurrency

type WebsiteChecker func(string) bool

type result struct {
	Str  string
	Bool bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChan := make(chan result)

	for _, url := range urls {
		go func(url string) {
			resultChan <- result{Str: url, Bool: wc(url)}
		}(url)
	}

	for i := 0; i < len(urls); i += 1 {
		result := <-resultChan
		results[result.Str] = result.Bool
	}

	return results
}
