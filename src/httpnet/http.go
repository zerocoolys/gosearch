package httpnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"request"
)

func Send(query request.Query) {
	var bytes, _ = json.Marshal(query)

	fmt.Println("request ", string(bytes))
	var resp = newHttpRequest(bytes)

	fmt.Println(resp)
}

func newHttpRequest(body []byte) (response *http.Response) {

	var reader = bytes.NewReader(body)
	var req, err = http.NewRequest("POST", "http://127.0.0.1:9200", reader)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("--> http request : ", req)
	var client = http.DefaultTransport
	var resp *http.Response
	resp, err = client.RoundTrip(req)
	if err != nil {
		return nil
	}

	return resp

}
