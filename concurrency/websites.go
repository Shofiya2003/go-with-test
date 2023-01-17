package concurrency

type WebsiteChecker func(website string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, websites []string) map[string]bool {
	res := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range websites {
		go func(url string) {
			resultChannel <- result{url, wc(url)}
		}(url)
	}

	for i := 0; i < len(websites); i++ {
		r := <-resultChannel
		res[r.string] = r.bool
	}

	return res
}
