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
	
	sender := siplib.NewSIPRecipient("sipfurious", "2000", "1.3.3.7", 5060)
	receiver := siplib.NewSIPRecipient("Joe Jones", "user2", "192.168.1.8", 5060)
	res,err := siplib.MapOptionsUDP("192.168.1.8", 5060, 10, sender, receiver)
	
	fmt.Println(res)
	fmt.Println(err)
}

func map_udp(targets []string, port int, timeout int) {
	res_targets := []string{}
	results := []string{}
	for _,target := range targets {
		fmt.Printf("Trying %s:%d...\n", target, port)
		result,err := siplib.MapUDP(target, port, timeout)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not map %s:%d (%s)\n", target, port, err.Error())
		} else {
			res_targets = append(res_targets, target)
			results = append(results, result)
		}
	}
	fmt.Println("")
	if len(res_targets) > 0 {
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 2, '\t', 0)
		fmt.Fprintf(w, "Target\tPort\tServer Header\n")
		fmt.Fprintf(w, "\t\t\t\n")
		for index,_ := range res_targets {
			fmt.Fprintf(w, "%s\t%d\t%s\n", res_targets[index], port, results[index])
		}
		w.Flush()
	} else {
		fmt.Println("No results found.")
	}
}
