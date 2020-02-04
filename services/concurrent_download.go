package services

import (
	"io"
	"net/http"
	"os"
	"path"
)

func Download(URL string) error {
	resp, err := http.Get(URL)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	filepath := path.Base(resp.Request.URL.String())
	out, err := os.Create(filepath)

	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	defer out.Close()

	return nil
}

func ConcurrentDownload(urls []string) (*map[string]string, bool) {

	 MappedURL := make(map[string]string)

	 urlCount := 0
	 successCount := 0

	 for _, URL := range urls {
	 	urlCount++
	 	go func(URL string) {
			err := Download(URL)

			if err != nil {
				MappedURL[URL] = "Unsuccessful"
				successCount++
			} else {
				MappedURL[URL] = "Successful"
			}

		}(URL)
	 }


	if urlCount == len(urls) {
		if successCount == urlCount {
			return &MappedURL, true
		} else {
			return &MappedURL, false
		}
	}

	return nil, false
}
