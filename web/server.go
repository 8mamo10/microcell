package main

import (
	"fmt"
	"html"
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
	fmt.Fprintf(w, "startStreaming:\nError:\n%v\n", err)
}

func stopStreamingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/home/pi/git/microcell/streaming/stop-streaming.sh").Output()
	fmt.Fprintf(w, "stopStreaming:\nError:\n%v\n", err)
}

func startFeedingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/home/pi/git/microcell/feeding/start-feeding.sh").Output()
	fmt.Fprintf(w, "startFeeding:\nError:\n%v\n", err)
}

func stopFeedingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := exec.Command("/home/pi/git/microcell/feeding/stop-feeding.sh").Output()
	fmt.Fprintf(w, "stopFeeding:\nError:\n%v\n", err)
}

func main() {
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
	s.ListenAndServe()
}
