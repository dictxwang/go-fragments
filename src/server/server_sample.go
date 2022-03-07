package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

/*
 *created by wangqiang at 2018/12/19
 */

func sayHello(w http.ResponseWriter, r *http.Request) {
	if ok, _ := regexp.MatchString("/static/", r.URL.String()); ok {
		staticServer(w, r)
		return
	}
	io.WriteString(w, "hello word")
}

func sayBey(w http.ResponseWriter, r * http.Request) {
	io.WriteString(w, "Bye bye")
}

func staticServer(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	http.StripPrefix("/static/",
		http.FileServer(http.Dir(wd))).ServeHTTP(w, r)
}

func loggingHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func SampleMain() {

	fmt.Println("\n[server_sample]")

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", sayHello)
	serveMux.Handle("/byebye", loggingHandler(http.HandlerFunc(sayBey)))
	serveMux.HandleFunc("/bye", sayBey)
	serveMux.HandleFunc("/static", staticServer)

	server := http.Server {
		Addr:	":8080",
		Handler:	serveMux,
		ReadTimeout:	5 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}