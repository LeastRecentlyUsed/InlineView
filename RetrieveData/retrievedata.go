package retrievedata

import "net/http"

// GetFileFromUrl returns an http response
func GetFileFromUrl(url string) *http.Response {
	resp, err := http.Get(url)
	return resp
}
