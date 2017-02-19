package fragments

import (
	"net/http"
	"io"
	"log"
	"strings"
	"os"
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

func putRequestJson(url string, path string) {
	// read the file
	data, err := os.Open(path)
	if err != nil {
		//handle error
		log.Fatal(err)
	}
	putRequest(url, data)

}

func main()  {
	putRequest("http://google.com", strings.NewReader("any thing"))

	putRequestJson("http://yahoo.com", "/path/to/file.json")
}