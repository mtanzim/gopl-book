package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch reading: %v\n", err)
		// 	os.Exit(1)
		// }
		fmt.Println(resp.Status)
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		// fmt.Printf("%s", b)

	}
}
