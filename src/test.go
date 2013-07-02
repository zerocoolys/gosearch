package main

// import (
// 	//"encoding/json"
// 	"bytes"
// 	"flag"
// 	"fmt"
// 	"httpnet"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"request"
// 	// "strconv"
// 	// "strings"
// )

// var (
// 	host        = flag.String("host", "127.0.0.1", "target host address")
// 	queryString = flag.String("q", "", "query string")
// )

// func main2() {

// 	flag.Parse()
// 	var content []byte

// 	// var body = "Hello, world."

// 	// content = strconv.AppendQuote(content, body)
// 	var reader io.Reader = bytes.NewReader(content)

// 	fmt.Println(content)

// 	var req, _ = http.NewRequest("GET", "http://127.0.0.1:9200", reader)

// 	client := http.DefaultClient

// 	var resp, err = client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("body length: %d\n", int(resp.ContentLength))
// 	var body, _ = ioutil.ReadAll(resp.Body)
// 	fmt.Printf("body content: %s\n", string(body))

// }

// func main1() {
// 	var queryStringQuery = request.QueryStringQuery{"hello"}
// 	var queryStringQueryBuilder = request.QueryStringQueryBuilder{queryStringQuery}
// 	var query = request.Query{queryStringQueryBuilder}

// 	httpnet.Send(query)

// }

func main() {
	array := make([]int, 1)
	array[0] = 1
	i := len(array)
	array = make([]int, i+1)
	array[1] = 1

	maps := make(map[string]string, 10)
	maps["name"] = "nil"
}
