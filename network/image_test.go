package network

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_FetchImageContentLength(t *testing.T) {
	mockContentLength := rand.Int63()
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		// equals(t, req.URL.String(), "/some/path")
		// Send response to be tested
		rw.Header().Set("Content-Length", strconv.FormatInt(mockContentLength, 10))
		// rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	imageRequest := &ImageRequest{
		Client:  server.Client(),
		BaseURL: server.URL,
	}

	// base := "http://www.liuwill.com"
	imgUrl := "/resources/sku-da2afa8f-17f7-6d2d-2635-f30020b10d58.png"
	contentLen, err := imageRequest.FetchImageContentLength(imgUrl)

	if err != nil || contentLen != mockContentLength {
		t.Error("Test FetchImageContentLength Fail", contentLen, mockContentLength)
	}
}
