package siplib

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
