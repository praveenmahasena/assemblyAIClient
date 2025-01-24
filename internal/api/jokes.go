package api

import (
	"io"
	"net/http"
)

func JokesPls() ([]byte, error) {
	res, resErr := http.Get("https://storage.googleapis.com/aai-docs-samples/sports_injuries.mp3")
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()
	result, resultErr := io.ReadAll(res.Body)
	if resultErr != nil {
		return nil, resultErr
	}
	return result, nil
}
