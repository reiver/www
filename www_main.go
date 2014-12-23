package main


import "fmt"
import "io"
import "io/ioutil"
import "net/http"
import "log"


func www_main(port int, reader io.Reader) {

	// Read everything from the reader.
	//
	// (This reader could be STDIN (i.e., os.Stdin) or it could be a file or pipe
	// that was specified by name on the command line.)
	//
	// NOTE that if the user sends a never ending stream to this program or
	// just a very very large stream to this program, that may cause problems!
	//
		bytes, err := ioutil.ReadAll(reader)
		if nil != err {
			www_500(port)
		}

	// Register handler for HTTP server.
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write(bytes)
		})

	// Launch HTTP server.
		addr := fmt.Sprintf(":%d", port)
		log.Fatal(http.ListenAndServe(addr, nil))
}
