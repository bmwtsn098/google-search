package fuzz
import "golang.org/x/time/rate"
import "github.com/rocketlaunchr/google-search"

func mayhemit(bytes []byte) int {
    content := string(bytes)

	googlesearch.RateLimit = rate.NewLimiter(5, 0)
	opts := googlesearch.SearchOptions{
		Limit: 20,
	}

	googlesearch.Search(nil, content, opts)
	return 0
}

func Fuzz(data []byte) int {

	_ = mayhemit(data)
	return 0
}