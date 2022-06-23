package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

var BaseUrl string

func main() {
    setInit()
    towns := GetRegions()
    for _, t := range towns {
        fmt.Println(t)
    }
    // Insert(towns)

    router := mux.NewRouter()
    router.HandleFunc("/", callback)
    http.Handle("/", headerMiddleWare(router))
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Fprintf(os.Stderr, "http error", err)
    }
}

func setInit() {
    BaseUrl = "https://pokemon.fandom.com"
}

func callback(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var output string
    switch len(vars) {
    case 0:
        output = "Hello, World!"
    }
    w.Write([]byte(output))
}

func headerMiddleWare(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        h.ServeHTTP(w,r)
    })
}
