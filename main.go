package main


import "flag"
import "fmt"
import "os"


func main() {

	// Define a variable to store the TCP port number.
	//
	// NOTE that we must call flag.Parse() to actually get the value in there.
	// (That is done later in the code.)
	//
	// This TCP port number will be the port number we launch the HTTP server on.
	//
	// It defaults to TCP port number 5555. But the user can override it with
	// the --port switch.
	//
	// Example:
	//
	//     www --port=17273
	//
		port_ptr := flag.Int("port", 5555, "The TCP port to run the HTTP server on. (Maximum permitted value is 65535.)")

	// Define a variable to store the demo mode.
	//
	// Demo modes let users see parts of this program that they wouldn't normally be able to see.
	// (Such as being able to see what this program shows when a "500 Internal Server Error" happens,
	// without actually having to have that error happen.)
	//
	// The default is NOT to run this in demo mode.
	//
		demo_mode_ptr := flag.String("demo", "none", "The demo mode to run in. (Default is \"none\".) Option is: \"500\".")

	// The method flag.Parse() must be called before we can use and of the defined flags.
	//
	// (Such as "--port".)
	//
		flag.Parse()

	// NOTE that TCP ports are 16-bit (unsigned) integers. Therefore, the maxium value for a
	// TCP port is 65535.
	//
	// Also, the minimum port number is 0 (zero).
	//
	// Thus, we need to check to make sure that the user specified a port number in the (closed) interval:
	// [0, 65535]
	//
		if port := *port_ptr ; 65535 < port {
			fmt.Fprintf(os.Stderr, "ERROR: You cannot pick a port number above 65535. (You chose %d.) Try again.\n", port)
			os.Exit(1)
			return
		} else if port := *port_ptr ; 0 > port {
			fmt.Fprintf(os.Stderr, "ERROR: You cannot pick a port number below 0. (You chose %d.) Try again.\n", port)
			os.Exit(1)
			return
		}

	// Check to see if this program is receiving input from STDIN.
	//
	// I.e., check to see if data is being piped to this program.
	//
		have_input_on_stdin := false

		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			have_input_on_stdin = true
		} else {
			have_input_on_stdin = false
		}

	// If we are NOT receiving input from STDIN, then...
	// Check to see if this program should be receiving input from a file name or pipe name specified on the
	// command line.
	//
	// NOTE that we only check this if we do NOT have input on the STDIN.
	//
	// Therefore, if the user sends an input stream on STDIN **and** species a file name (or pipe name) via
	// the command line, because there was an input stream on STDIN the file name (or pipe name) specified
	// on the command line WILL BE IGNORED.
	//
	// NOTE that we need flag.Parse() to run before checking this.
	//
		have_input_from_command_line := false
		cmdin_name := ""

		args := flag.Args()
		if ! have_input_on_stdin && 1 <= len(args) {
			arg0 := args[0]

			if _, err := os.Stat(arg0); ! os.IsNotExist(err) {
				have_input_from_command_line = true
				cmdin_name = arg0
			}
		}

	// Send a message to the user telling them how they can connect to the HTTP server.
		fmt.Printf("\nView locally at URL:\nhttp://127.0.0.1:%d/\n\nPress [CTRL]+[C] to quit.\n\n", *port_ptr)

	// Launch one the appropriate HTTP server.
	//
	// Here we have a different HTTP server depending on whether the programing was sent something on STDIN or not.
	//
	// And depending on whether this is being run in demo mode of not.
	//
		switch *demo_mode_ptr {
			case "500":
				www_500(*port_ptr)
			default:
				switch have_input_on_stdin {
					case true:
						www_main(*port_ptr, os.Stdin)
					case false:
						switch have_input_from_command_line {
							case true:
								cmdin, err := os.Open(cmdin_name)
								if nil != err {
									www_500(*port_ptr)
								} else {
									www_main(*port_ptr, cmdin)
								}
							case false:
								www_help(*port_ptr)
						}
				}
		}
}
