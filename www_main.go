package main


import "fmt"
import "io/ioutil"
import "net/http"
import "log"
import "os"


func www_main(port int) {

	// Read everything from STDIN.
	//
	// NOTE that if the user sends a never ending stream to this program or
	// just a very very large stream to this program, that may cause problems!
	//
		bytes, err := ioutil.ReadAll(os.Stdin)
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
