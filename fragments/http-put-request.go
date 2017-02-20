package fragments

import (
	"net/http"
	"io"
	"log"
	"strings"
	"os"
	"bytes"
)

func putRequest(url string, data io.Reader)  {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	_, err = client.Do(req)
	if err != nil {
		// handle error
		log.Fatal(err)
	}


}

func httpPutExample()  {
	putRequest("http://google.com", strings.NewReader("any thing"))

	var jsonStr string = []bytes {"name":"Rob", "title":"developer"}
	putRequest("http://msn.com", bytes.NewBuffer(jsonStr))

	// read the file
	data, err := os.Open("/path/to/file.json")
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	putRequest("http://yahoo.com", data)
}