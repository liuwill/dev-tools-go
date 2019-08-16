package network

import (
	"errors"
	"net/http"
	"strconv"
)

type ImageRequest struct {
	Client  *http.Client
	BaseURL string
}

func (request *ImageRequest) FetchImageContentLength(url string) (int64, error) {
	resp, err := request.Client.Head(request.BaseURL + url) //发送请求

	if err != nil || resp.StatusCode != http.StatusOK {
		return 0, errors.New("request fail")
	}
	defer resp.Body.Close()

	contentLength := resp.Header["Content-Length"][0]
	// fmt.Println("rs.Header:", contentLength)
	lenVal, _ := strconv.ParseInt(contentLength, 10, 64)
	// if err != nil {
	// 	return 0, err
	// }
	return lenVal, nil
}
