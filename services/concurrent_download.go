package services

func ConcurrentDownload(urls []string) *map[string]string {

	 MappedURLs := make(map[string]string)

	 urlCount := 0

	 for _, URL := range urls {
	 	urlCount++
	 	go func(URL string) {
			err := Download(URL)

			if err == true {
				MappedURLs[URL] = "Unsuccessful"
			} else {
				MappedURLs[URL] = "Successful"
			}

		}(URL)
	 }

	if urlCount == len(urls) {
		return &MappedURLs
	} else {
		return nil
	}
}
