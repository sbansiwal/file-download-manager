package services

import (
	"io"
	"net/http"
	"os"
	"path"
)

func SerialDownload(urls []string) (*map[string]string, bool) {

	MappedUrls := make(map[string]string)

	isSuccess := true

	for _, URL := range urls {
		resp, err := http.Get(URL)

		if err != nil {
			MappedUrls[URL] = "Unsuccessful"
			isSuccess = false
			continue
		}

		filepath := path.Base(resp.Request.URL.String())
		out, err := os.Create(filepath)

		if err != nil {
			MappedUrls[URL] = "Unsuccessful"
			isSuccess = false
			continue
		}

		_, err = io.Copy(out, resp.Body)

		if err != nil {
			MappedUrls[URL] = "Unsuccessful"
			isSuccess = false
			continue
		}

		MappedUrls[URL] = "Successful"

		_ = resp.Body.Close()

		_= out.Close()

	}

	return &MappedUrls, isSuccess
}
