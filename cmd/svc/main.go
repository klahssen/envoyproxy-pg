package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-zoo/bone"
)

var svc = os.Getenv("SERVICE")

func main() {
	port := "8080"
	pport := flag.String("p", "", "port on which the server will listen")
	flag.Parse()
	if *pport == "" {
		if p := os.Getenv("SERVER_PORT"); p != "" {
			//fmt.Printf("using SERVER_PORT environment variable\n")
			port = p
		} else {
			//fmt.Printf("using default PORT\n")
		}
	} else {
		//fmt.Printf("using PORT from -p flag\n")
		port = *pport
	}
	p := 0
	var err error
	if p, err = strconv.Atoi(port); err != nil {
		panic(fmt.Sprintf("invalid port '%s': %s", port, err.Error()))
	}
	server := http.Server{
		Addr:              fmt.Sprintf(":%d", p),
		ReadTimeout:       time.Second * 3,
		WriteTimeout:      time.Second * 3,
		ReadHeaderTimeout: time.Second * 3,
		IdleTimeout:       time.Second * 3,
		Handler:           getMux(),
	}
	log.Printf("Listening on port %d", p)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func getMux() http.Handler {
	mux := bone.New()
	mux.GetFunc(fmt.Sprintf("/%s", svc), hello)
	mux.GetFunc(fmt.Sprintf("/%s/:name", svc), helloName)
	return logMiddleware(mux)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Printf("received a request to /hello")
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	name := r.URL.Query().Get("name")
	str := ""
	if name == "" {
		str = "hello from " + svc + "\n"
	} else {
		str = fmt.Sprintf("hello '%s' from svc '%s'\n", name, svc)
	}
	fmt.Fprintf(w, str)
}
func helloName(w http.ResponseWriter, r *http.Request) {
	name := bone.GetValue(r, "name")
	fmt.Fprintf(w, fmt.Sprintf("hello '%s' from svc '%s'\n", name, svc))
}

func logMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
