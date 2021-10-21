package main
  
import (
        "fmt"
        "log"
        "net/http"
        "time"
)

func main() {
        http.HandleFunc("/", Use)
        http.HandleFunc("/time", serveTime)
        http.HandleFunc("/nanotime", serveNanoTime)
        log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func Use(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Use:  <URL>/time  or  <URL>/nanotime")
}

func serveTime(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Current time: ", time.Now().String())
}

func serveNanoTime(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "UnixNano time: ", time.Now().UnixNano())
}

