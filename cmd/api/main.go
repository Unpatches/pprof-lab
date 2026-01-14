package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // регистрирует /debug/pprof/*

	"example.com/pprof-lab/internal/work"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/debug/pprof/", http.DefaultServeMux)
	mux.Handle("/debug/pprof", http.DefaultServeMux)

	// Эндпоинт, вызывающий “тяжёлую” работу.
	mux.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		defer work.TimeIt("Fib(38)")()
		res := work.Fib(38)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte((fmtInt(res))))
	})

	mux.HandleFunc("/work-fast", func(w http.ResponseWriter, r *http.Request) {
		defer work.TimeIt("FibFast(38)")()
		res := work.FibFast(38)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte((fmtInt(res))))
	})

	log.Println("Server on :8080; pprof on /debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func fmtInt(v int) string { return fmt.Sprintf("%d\n", v) }
