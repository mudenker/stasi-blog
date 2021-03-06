package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serve(directoryToServe, basepath string, port int) {
	//Example in my case.
	//go run . --input="../blog-test-source" --output="../blog-test" & go run demo/server.go --dir="../blog-test" --basepath="/blog-test/"

	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
		os.Exit(0)
	}()

	log.Printf("Serving %s at localhost:%d/%s", directoryToServe, port, basepath)
	log.Println("Please remember to only use this command to serve your website in a development scenario.")
	portString := fmt.Sprintf(":%d", port)

	dir := dirWith404Handler{http.Dir(directoryToServe)}
	if basepath == "" {
		log.Fatal(http.ListenAndServe(portString, http.FileServer(dir)))
	} else {
		http.Handle(basepath, http.StripPrefix(basepath, http.FileServer(dir)))
		log.Fatal(http.ListenAndServe(portString, nil))
	}
}

type dirWith404Handler struct {
	dir http.Dir
}

// Open implements FileSystem using os.Open, opening files for reading rooted
// and relative to the directory d. If a file can't be found, we return a 404
// page instead.
func (d dirWith404Handler) Open(name string) (http.File, error) {
	file, err := d.dir.Open(name)
	if os.IsNotExist(err) {
		file404, newError := d.dir.Open("404.html")
		if newError != nil {
			return nil, newError
		}
		//Technically we'd need the old error to indicate 404 to the
		//browser, but for demo/test purposes, this'll do.
		return file404, nil
	}
	return file, err
}
