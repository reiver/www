package main


import "fmt"
import "net/http"
import "log"


func www_500(port int) {

	// Construct the help message.
		html := `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8" />
		<title>500 Internal Server Error | www</title>
	</head>

	<body>
		<h1>Ooops! Something went wrong!</h1>
		<p>
			<span style="font-size:10em;">&#9760;</span>
		</p>
		<p>
			If you are seeing this message you may have tried sending something
			that was way too big for this <b>www</b> command and your computer to handle.
		</p>

		<footer style="margin-top:7em;">
			<hr />
			<p>
				To visit the <b>www</b> project homepage, go to:
				<br />
				<a href="https://github.com/reiver/www">github.com/reiver/www</a>
			</p>
			<span style="font-size:2em;">✪</span>
		</footer>
	</body>

</html>
`

	// Register handler for HTTP server.
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, html)
		})

	// Launch HTTP server.
		addr := fmt.Sprintf(":%d", port)
		log.Fatal(http.ListenAndServe(addr, nil))
}
