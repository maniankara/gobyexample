package fragments

import (
	"net/http"
	"bytes"
)

// define a type function
type httpReqFunc func () (resp *http.Response, err error)

func example(fn httpReqFunc, ret int) {
	resp, err := fn()

	// perform some actions/handling here

	return resp, err
}

func caller()  {

	example(func() (resp *http.Response, err error) {
		return http.Post("http://google.com", "application/json", bytes.NewReader(""))
	}, 10)
}


