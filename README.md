# aero [![Go Report Card](https://goreportcard.com/badge/github.com/aldidana/aero)](https://goreportcard.com/report/github.com/aldidana/aero)

Aero is tiny Express-inspired router for golang.

## Install
You need go version 1.7++
```sh
go get github.com/aldidana/aero
```

## Example
```go
package main

import (
	"fmt"
	"github.com/aldidana/aero"
	"net/http"
)

func main() {
	router := aero.Router()

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Index Page"))
	})

	// Using params
	// This handler will match /hello/bingo but it will not match /hello/ and /hello
	router.Get("/hello/:name", helloHandler)

	http.ListenAndServe(":5678", router)
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	name := req.Context().Value("name")

	fmt.Fprintf(res, "Hello %v\n", name)
}
```
## TODO
- [ ] Testing
- [ ] Group several routes
- [ ] Maybe route logging :beers:

## License
MIT @Aldi Priya Perdana