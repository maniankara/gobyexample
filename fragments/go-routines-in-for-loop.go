package fragments

func downloadUrls(urls []string) {

	done := make(chan bool, len(urls))

	for _, url := range urls {

		go func(url string) {

			// perform download here
			done <- true
		}(url)

	}
	// ensure all the routines are done
	for i:=0; i<len(urls); i++ { <- done }
}
