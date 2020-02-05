package services

func SerialDownload(urls []string) (*map[string]string, bool) {

	MappedUrls := make(map[string]string)

	isSuccess := true

	for _, URL := range urls {
		err := Download(URL)

		if err == true {
			MappedUrls[URL] = "Unsuccessful"
		} else {
			MappedUrls[URL] = "Successful"
		}
	}

	return &MappedUrls, isSuccess
}
