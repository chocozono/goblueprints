package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/chocozono/trace"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	//2nd argument is required for template
	t.templ.Execute(w, r)
}

func main() {
	//e.x. ./chat -addr=":3000"
	var addr = flag.String("addr", ":8080", "address for the application")
	flag.Parse()
	room := newRoom()
	room.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", room)

	// start chat room
	go room.run()

	//start webserver
	log.Println("Start Web server. port: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
