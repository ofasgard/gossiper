package main

import "siplib"
import "time"
import "math/rand"
import "fmt"

func main() {
	simpletest()
}

func simpletest() {
	rand.Seed(time.Now().UnixNano())
	
	sender := siplib.NewSIPRecipient("sipfurious", "2000", "1.3.3.7", 5060)
	receiver := siplib.NewSIPRecipient("Joe Jones", "user2", "192.168.1.8", 5060)
	req := siplib.NewOptionsRequest("UDP", sender, receiver)
	
	resp,err := siplib.RequestUDP("192.168.1.8", 5060, 10, req)
	fmt.Println(resp)
	fmt.Println(err)
	
}

// http://siptutorial.net/SIP/request.html
// https://datatracker.ietf.org/doc/html/rfc3261
