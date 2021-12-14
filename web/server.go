package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os/exec"
	"text/template"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("streaming/nginx/html/index.html")
	if err != nil {
		log.Printf("[index] Error:%v", err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Printf("[index] Error:%v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func psHandler(w http.ResponseWriter, r *http.Request) {
	ps, err := exec.Command("sh", "-c", "ps aux").CombinedOutput()
	fmt.Fprintf(w, "ps:\n%s :Error:\n%v\n", ps, err)
}

func startStreamingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("sh", "-c", "/var/www/bin/start-streaming.sh").CombinedOutput()
	log.Printf("[startStreaming] Error:%v", err)
	time.Sleep(5 * time.Second)
	http.Redirect(w, r, "/", 302)
}

func stopStreamingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("sh", "-c", "/var/www/bin/stop-streaming.sh").CombinedOutput()
	log.Printf("[stopStreaming] Error:%v", err)
	http.Redirect(w, r, "/", 302)
}

func startFeedingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("sh", "-c", "/var/www/bin/start-feeding.sh").CombinedOutput()
	log.Printf("[startFeeding] Error:%v", err)
	http.Redirect(w, r, "/", 302)
}

func stopFeedingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("sh", "-c", "/var/www/bin/stop-feeding.sh").CombinedOutput()
	log.Printf("[stopFeeding] Error:%v", err)
	http.Redirect(w, r, "/", 302)
}

func main() {
	log.Printf("Start\n")
	mux := http.NewServeMux()
	mux.HandleFunc("/index", indexHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/ps", psHandler)
	mux.HandleFunc("/startStreaming", startStreamingHandler)
	mux.HandleFunc("/stopStreaming", stopStreamingHandler)
	mux.HandleFunc("/startFeeding", startFeedingHandler)
	mux.HandleFunc("/stopFeeding", stopFeedingHandler)

	s := &http.Server{
		Addr:    ":55556",
		Handler: mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
