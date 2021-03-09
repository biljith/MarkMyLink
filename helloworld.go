package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", handle)
    http.HandleFunc("/hello", greet)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("Listening on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}

func handle(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    fmt.Fprint(w,  `<!DOCTYPE html>
                    <html lang="en">
                    <form action=hello>
                      <label for="fname">First name:</label><br>
                      <input type="text" id="fname" name="fname" value="John"><br>
                      <label for="lname">Last name:</label><br>
                      <input type="text" id="lname" name="lname" value="Doe"><br><br>
                      <input type="submit" value="Submit">
                    </form>`)
}

func greet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s %s", r.URL.Query().Get("fname"), r.URL.Query().Get("lname"))
}
