package main

import (
	"fmt"
	"html"
	"net/http"
	"os/exec"
	"text/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(w, nil); err != nil {
		panic(err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func psHandler(w http.ResponseWriter, r *http.Request) {
	ls, err := exec.Command("ps").Output()
	fmt.Fprintf(w, "ps:\n%s :Error:\n%v\n", ls, err)
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
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/ps", psHandler)
	http.HandleFunc("/startStreaming", startStreamingHandler)
	http.HandleFunc("/stopStreaming", stopStreamingHandler)
	http.HandleFunc("/startFeeding", startFeedingHandler)
	http.HandleFunc("/stopFeeding", stopFeedingHandler)
	http.ListenAndServe("192.168.86.111:6624", nil)
}
