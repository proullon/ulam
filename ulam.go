package main

import (
    "os"
    "log"
    "net/http"
    "path"
    "github.com/proullon/ulam/prime"
    "encoding/json"
    "fmt"
    "strconv"
)

type Ulam struct {

}

func (u *Ulam) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    // Parse size from url    r.ParseForm()
    r.ParseForm()
    size, err := strconv.Atoi(r.FormValue("size"))
    if err != nil {
        fmt.Fprintf(w, "%s", err)
        return
    }

    spiral, _ := prime.Ulam(int64(size))
    data, _ := json.Marshal(spiral)
    fmt.Fprintf(w, "%s", string(data))
}

func main() {
    static_directory, err := os.Getwd()
    if err != nil {
        log.Fatal("Cannot get working directory")
    }

    static_directory = path.Join(static_directory, "static")

    log.Printf("Launching Ulam\n")
    http.Handle("/api/ulam", &Ulam{})
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(static_directory))))
    err = http.ListenAndServe("127.0.0.1:8080", nil)
    if err != nil {
        log.Fatal("Cannot start http server : ", err)
    }
}
