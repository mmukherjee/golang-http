
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><head><title>Greetings to the world</title></head><body><h1>Hello from the Jenkins X golang http example. This is a new feature. But this also has to be there</h1></body></html>\n")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
