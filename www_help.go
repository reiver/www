package main


import "fmt"
import "net/http"
import "log"


func www_help(port int) {

	// Construct the help message.
		html := `<html>
	<head>
		<title>www</title>
	<head>

	<body>
		<h1>www</h1>
		<p>
			Thank you for using <b>www</b>.
		</p>
		<p>
			If you are seeing this message you may have forgotten to pipe something to
			this command.
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
