// serverを起動するmain関数を書く
package main

import (
	"net/http"
)

func main() {
	r := HandleRequests()
	http.ListenAndServe(":8080", r)
}