package services

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func DownloadErr(URL string) bool {
	resp, err := http.Get(URL)
	readRes, _ := http.Get(URL)

	if err != nil {
		return true
	}

	bodyBytes, err := ioutil.ReadAll(readRes.Body)

	validImage := CheckValidity(bodyBytes)

	if validImage == false {
		return true
	}


	filepath := path.Base(resp.Request.URL.String())
	out, err := os.Create(filepath)

	if err != nil {
		return true
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return true
	}

	defer resp.Body.Close()

	defer out.Close()

	return false
}
