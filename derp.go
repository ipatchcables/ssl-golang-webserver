package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.js"))
}
func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServeTLS("10.9.1.24:443","server.crt", "client.key", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
  tpl.ExecuteTemplate(w, "derp.js", res)
	ua := r.Header.Get("User-Agent")
	url := fmt.Sprintf("%v %v %v %v %v", r.Method, r.URL, r.Proto, r.Host, ua)
	fmt.Printf("Tracking: %s \n", url)
	return

}
