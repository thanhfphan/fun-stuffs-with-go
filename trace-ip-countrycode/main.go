package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://1.1.1.1/cdn-cgi/trace", nil)
	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(body), "\n")

	m := map[string]string{}
	for _, l := range lines {
		tmp := strings.Split(l, "=")
		if len(tmp) != 2 {
			continue
		}
		m[tmp[0]] = tmp[1]
	}

	fmt.Println("IP Address: ", m["ip"])
	fmt.Println("Location: ", m["loc"])

}
