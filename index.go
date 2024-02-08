package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/aarol/reload"
)

func main() {
	isDevelopment := flag.Bool("dev", true, "Development mode")
	port := flag.String("port", "3000", "Port to serve the app")
	flag.Parse()

	// server the public folder as static
	http.Handle("/", http.FileServer(http.Dir("public/")))

	// add api for htmx components
	http.Handle("/api/components/", http.StripPrefix("/api/components/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if *isDevelopment {
			// Use the request path as needed
			log.Println("Request path:", r.URL.Path)
			// yeet the cache
			w.Header().Set("Cache-Control", "no-store")
		}

		// server the static html file through the api
		http.ServeFile(w, r, "public/components/"+r.URL.Path+".html")
	})))

	// hot reload from aarol/reload
	var handler http.Handler = http.DefaultServeMux
	if *isDevelopment {
		// Call `New()` with a list of directories to recursively watch
		reloader := reload.New("public/")

		// Optionally, define a callback to
		// invalidate any caches
		// reloader.OnReload = func() {
		// parse templates, process scss, etc
		// }

		handler = reloader.Handle(handler)
		log.Println("Hot Reload Enabled...")
	}

	log.Println("Listening on port " + *port)
	err := http.ListenAndServe(":"+*port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
