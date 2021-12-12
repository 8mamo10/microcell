package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os/exec"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func psHandler(w http.ResponseWriter, r *http.Request) {
	ps, err := exec.Command("sh", "-c", "ps aux").CombinedOutput()
	fmt.Fprintf(w, "ps:\n%s :Error:\n%v\n", ps, err)
}

func startStreamingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/home/pi/git/microcell/streaming/start-streaming.sh").Output()
	log.Printf("[startStreaming] Error:%v", err)
	http.Redirect(w, r, "http://192.168.86.111:55555/", 302)
}

func stopStreamingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/home/pi/git/microcell/streaming/stop-streaming.sh").Output()
	log.Printf("[stopStreaming] Error:%v", err)
	http.Redirect(w, r, "http://192.168.86.111:55555/", 302)
}

func startFeedingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/home/pi/git/microcell/feeding/start-feeding.sh").Output()
	log.Printf("[startFeeding] Error:%v", err)
	http.Redirect(w, r, "http://192.168.86.111:55555/", 302)
}

func stopFeedingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/home/pi/git/microcell/feeding/stop-feeding.sh").Output()
	log.Printf("[stopFeeding] Error:%v", err)
	http.Redirect(w, r, "http://192.168.86.111:55555/", 302)
}

func main() {
	log.Printf("Start\n")
	mux := http.NewServeMux()
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
