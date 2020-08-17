package main

import (
	"log"
	"net/http"
	"time"
	_ "net/http/pprof"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
	time.Sleep(10 * time.Millisecond)
}

func main() {
	/*
		* HOW TO PROFILE GO APP?
		* install https://github.com/google/pprof locally (also graphviz for UI)
		* expose the 6060 port in docker and then go http://localhost:6060/debug/pprof
		* trigger one of the endpoints, e.g: http://localhost:6060/debug/pprof/profile, this will
	download "profile" file after 60 seconds
		* while waiting 60 seconds, also run a load test, e.g: ab -n 1010 -c 10 http://localhost:8002/hello
		* run following tool by showing downloaded file, in this example "profile" is the file: go tool pprof -http=":8080" profile
		*
		* Detailed blog post: https://bruinsslot.jp/post/profiling-golang-docker-2/
		*
	 */
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	s := &server{}
	http.Handle("/hello", s)
	// Because Dispacther adds / if it missing for some reason
	http.Handle("/hello/", s)
	log.Fatal(http.ListenAndServe(":8000", nil))
}