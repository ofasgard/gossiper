package main

import "siplib"
import "time"
import "math/rand"
import "os"
import "fmt"
import "flag"
import "text/tabwriter"

// For SIP implementation info see:
// http://siptutorial.net/SIP/request.html
// https://datatracker.ietf.org/doc/html/rfc3261

func main() {
	rand.Seed(time.Now().UnixNano())
	//parse flags
	flag.Usage = usage
	port_ptr := flag.Int("port", 5060, "")
	timeout_ptr := flag.Int("timeout", 10, "")
	threads_ptr := flag.Int("threads", 5, "")
	flag.Parse()
	timeout := *timeout_ptr
	threads := *threads_ptr
	port := *port_ptr
	//validate args
	if flag.NArg() < 3 {
		usage()
		return
	}
	method := flag.Arg(0)
	protocol := flag.Arg(1)
	targets := parse_target(flag.Arg(2))

	//temporary test stuff	
	protocol = method
	method = protocol
	results := siplib.ScanOptionsUDP(targets, port, timeout, threads)
	for target, result := range results {
		fmt.Printf("%s: %s\n", target, result)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <map|war|crack> <udp|tcp|tls> <target>\n\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "'map': Scanner that uses OPTIONS to attempt to retrieve the SIP Server header.\n")
	fmt.Fprintf(os.Stderr, "'war': Wardialler that bruteforces extensions using the INVITE method.\n")
	fmt.Fprintf(os.Stderr, "'crack': Bruteforcer to crack SIP passwords for an extension.\n")
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Optional arguments:\n")
	w := new(tabwriter.Writer)
	w.Init(os.Stderr, 0, 8, 2, '\t', 0)
	fmt.Fprintf(w, "\t--port <#>\tPort to connect to SIP servers on. [DEFAULT: 5060]\n")
	fmt.Fprintf(w, "\t--timeout <sec>\tTimeout (in seconds) for each request. [DEFAULT: 10]\n")
	fmt.Fprintf(w, "\t--threads <#>\tNumber of hosts to target simultaneously. [DEFAULT: 5]\n")
	w.Flush()
	fmt.Fprintf(os.Stderr, "\n\nExample: %s map udp 192.168.0.20\n", os.Args[0])
}
