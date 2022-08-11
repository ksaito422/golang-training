package main

import (
	"fmt"
	"io"
	"net/http"
)

// HTTPErrorはステータスコードが200以外の場合のエラーを扱う構造体
type HTTPError struct {
	StatusCode int
	URL        string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", he.StatusCode, he.URL)
}

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのステータスコードが200ではない場合は *HTTPErrorとしてエラーを返却
	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}
	return io.ReadAll(resp.Body)
}
