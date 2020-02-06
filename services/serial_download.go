package services

func SerialDownload(urls []string) (*map[string]string, bool) {

	MappedUrls := make(map[string]string)

	isSuccess := true

	for _, URL := range urls {
		err := DownloadErr(URL)

		if err == true {
			MappedUrls[URL] = "Unsuccessful"
			isSuccess = false
		} else {
			MappedUrls[URL] = "Successful"
		}
	}

	return &MappedUrls, isSuccess
}
