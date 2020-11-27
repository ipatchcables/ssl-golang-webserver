package main

import (
        "fmt"
        "github.com/gorilla/mux"
        "log"
        "net/http"
        "net/http/httputil"
          "html/template"
)
var tpl * template.Template


func init() {
    tpl = template.Must(template.ParseGlob("templates/*.js"))
}

func DumpRequest(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/javascript")
        tpl.ExecuteTemplate(w, "derp.js", r)
        requestDump, err := httputil.DumpRequest(r, true)
        if err != nil {
                fmt.Println(w, err.Error())
        } else {
                fmt.Println(w, string(requestDump))
        }
}

func main() {
        router := mux.NewRouter()
        router.HandleFunc("/", DumpRequest).Methods("GET")
        router.HandleFunc("/", DumpRequest).Methods("POST")
        log.Fatal(http.ListenAndServeTLS("10.9.1.24:443", "server.crt", "server.key",router))
}
