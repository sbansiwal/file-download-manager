package services

func ConcurrentDownload(urls []string) (*map[string]string, bool){

	 MappedURLs := make(map[string]string)
	isSuccess := true
	 urlCount := 0

	 for _, URL := range urls {
	 	urlCount++
	 	go func(URL string) {
			err := DownloadErr(URL)

			if err == true {
				MappedURLs[URL] = "Unsuccessful"
				isSuccess = false
			} else {
				MappedURLs[URL] = "Successful"
			}

		}(URL)
	 }

	if urlCount == len(urls) {
		return &MappedURLs, isSuccess
	} else {
		return nil, false
	}
}
