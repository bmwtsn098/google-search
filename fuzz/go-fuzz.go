package fuzz
import "github.com/rocketlaunchr/google-search"

func mayhemit(bytes []byte) int {
    content := string(bytes)

	var RateLimit = rate.NewLimiter(5, 0)
	opts := googlesearch.SearchOptions{
		Limit: 20,
	}

	returnLinks, err := googlesearch.Search(nil, content, opts)
	return 0
}

func Fuzz(data []byte) int {

	_ = mayhemit(data)
	return 0
}