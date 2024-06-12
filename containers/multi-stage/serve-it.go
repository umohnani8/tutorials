package main

import (
    "fmt"
    "net/http"
)

const asciiArt = `
<pre>
  _____
 /     \
| () () |
 \  ^  /
  |||||
  |||||
</pre>
`

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><body><h1>Welcome to the Go ASCII Art Server!</h1>%s</body></html>", asciiArt)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server at :8080")
    http.ListenAndServe(":8080", nil)
}
