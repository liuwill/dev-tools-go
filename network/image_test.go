package network

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_FetchImageContentLength(t *testing.T) {
	errorUrl := "/error/do"
	mockContentLength := rand.Int63()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		// Send response to be tested
		if req.URL.String() == errorUrl {
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			rw.Header().Set("Content-Length", strconv.FormatInt(mockContentLength, 10))
			rw.WriteHeader(http.StatusOK)
		}
		// rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	imageRequest := &ImageRequest{
		Client:  server.Client(),
		BaseURL: server.URL,
	}

	// base := "http://www.liuwill.com"
	imgUrl := "/resources/sku-da2afa8f-17f7-6d2d-2635.png"
	contentLen, err := imageRequest.FetchImageContentLength(imgUrl)

	if err != nil || contentLen != mockContentLength {
		t.Error("Test FetchImageContentLength Fail", contentLen, mockContentLength)
	}

	_, err = imageRequest.FetchImageContentLength(errorUrl)
	if err == nil {
		t.Error("Test FetchImageContentLength Error")
	}
}
