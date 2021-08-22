package main

import "siplib"
import "fmt"
import "time"
import "math/rand"

func main() {
	rand.Seed(time.Now().UnixNano())
	sender := siplib.NewSIPRecipient("sipfabulous", "2000", "192.168.1.220", 5060)
	receiver := siplib.NewSIPRecipient("sipferrous", "user2", "server.com", 5060)
	
	x := siplib.NewOptionsRequest("UDP", sender, receiver)
	fmt.Println(x.Generate())
}

// http://siptutorial.net/SIP/request.html
// https://datatracker.ietf.org/doc/html/rfc3261
