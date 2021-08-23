package siplib

import "fmt"

// Single-threaded OPTIONS scan via UDP.

func MapOptionsUDP(target string, port int, timeout int, sender SIPRecipient, receiver SIPRecipient) (string,error) {
	req := NewOptionsRequest("UDP", sender, receiver)
	res,err := RequestUDP(target, port, timeout, req)
	if err != nil {
		return "",err
	}
	if val,ok := res.Headers["Server"]; ok {
		return val,nil
	}
	if val,ok := res.Headers["User-Agent"]; ok {
		return val,nil
	}
	return "[NONE]",nil
}

// Single-threaded OPTIONS scan via TCP.

func MapOptionsTCP(target string, port int, timeout int, sender SIPRecipient, receiver SIPRecipient) (string, error) {
	req := NewOptionsRequest("TCP", sender, receiver)
	res,err := RequestTCP(target, port, timeout, req)
	if err != nil {
		return "",err
	}
	if val,ok := res.Headers["Server"]; ok {
		return val,nil
	}
	if val,ok := res.Headers["User-Agent"]; ok {
		return val,nil
	}
	return "[NONE]",nil
}

// Multi-threaded OPTIONS scan via UDP.

func ScanOptionsUDP(targets []string, port int, timeout int, threads int) {
	// Anonymous function for workers
	worker := func(input chan string, output chan string, port int, timeout int) {
		jobs := []string{}
		//Get input from the channel until it is closed.
		for {
			val,ok := <-input
			if !ok { break }
			jobs = append(jobs, val)
		}
		//Perform an options scan on each target; send results to the output channel.
		for _,job := range jobs {
			sender := NewSIPRecipient("sipfurious", "100", "1.1.1.1", 5060)
			receiver := NewSIPRecipient("sipfurious", "200", job, port)
			res,err := MapOptionsUDP(job, port, timeout, sender, receiver)
			if err == nil {
				output <- res
			}
		}
		//Close the output channel.
		close(output)
	}
	
	//Build a worker pool equal to the desired thread limit.
	input_channels := []chan string{}
	output_channels := []chan string{}
	for i := 0; i<threads; i++ {
		input := make(chan string, 0)
		output := make(chan string, 0)
		go worker(input, output, port, timeout)
		input_channels = append(input_channels, input)
		output_channels = append(output_channels, output)
	}
	
	//Allocate jobs to the workers.
	current_worker := 0
	for i := 0; i<len(targets); i++ {
		input_channels[current_worker] <- targets[i]
		current_worker++
		if current_worker == len(input_channels) {
			current_worker = 0
		}
	}
	//Close input channels.
	for _,channel := range input_channels {
		close(channel)
	}
	//Wait for output channels to return.
	output := make(map[string]string)
	for _,channel := range output_channels {
		for {
			val,ok := <-channel
			if !ok { break }
			output["unknown"] = val
		}
	}
	fmt.Println(output)
	
	//needs a way to correlate output with targets...
}
