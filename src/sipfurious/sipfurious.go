package main

import "siplib"
import "time"
import "math/rand"
import "fmt"

// For SIP implementation info see:
// http://siptutorial.net/SIP/request.html
// https://datatracker.ietf.org/doc/html/rfc3261

func main() {
	simpletest()
}

func simpletest() {
	rand.Seed(time.Now().UnixNano())
	
	results := siplib.ScanOptionsUDP([]string{"192.168.1.8", "192.168.1.8", "192.168.1.8"}, 5060, 10, 5)
	for target, result := range results {
		fmt.Printf("%s: %s\n", target, result)
	}

}
