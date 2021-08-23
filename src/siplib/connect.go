package siplib

import "fmt"
import "net"
import "time"

// Create a UDP connection using source port 5060.
// If we can't nab it, use a random source port.

func ConnectUDP(target_host string, target_port int) (*net.UDPConn,error) {
	target := fmt.Sprintf("%s:%d", target_host, target_port)
	target_addr,err := net.ResolveUDPAddr("udp4", target)
	if err != nil {
		return nil,err
	}
	local_addr,err := net.ResolveUDPAddr("udp", ":5060")
	if err != nil {
		return nil,err
	}
	conn,err := net.DialUDP("udp4", local_addr, target_addr)
	if err == nil {
		return conn,nil
	}
	conn,err = net.DialUDP("udp4", nil, target_addr)
	return conn,err
}

// Send a SIPRequest along an open UDP connection and return the response (or an error).
// You should use a goroutine or conn.SetDeadline() to ensure it doesn't block forever.

func SendUDP(conn *net.UDPConn, req SIPRequest) error {
	msg := req.Generate()
	_,err := fmt.Fprintf(conn, msg)
	return err
}

func RecvUDP(conn *net.UDPConn) (string,error) {
	output := ""
	for {
		buf := make([]byte, 8192)
		n,err := conn.Read(buf)
		if err != nil {
			return output,err
		}
		if n == 8192 {
			output += string(buf)
		} else {
			output += string(buf[0:n])
			break
		}
	}
	return output, nil
}

func RequestUDP(target string, port int, timeout int, req SIPRequest) (SIPResponse,error) {
	conn,err := ConnectUDP(target, port)
	if err != nil {
		return SIPResponse{},err
	}
	
	defer conn.Close()
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	conn.SetDeadline(deadline)
	
	err = SendUDP(conn, req)
	if err != nil {
		return SIPResponse{},err
	}
	
	//receive and parse responses until one matches or we time out
	call_id := ""
	if _, ok := req.Headers["Call-ID"]; ok {
		call_id = req.Headers["Call-ID"]
	}
	for {
		resp,err := RecvUDP(conn)
		if err != nil {
			return SIPResponse{},err
		}
		parsed,err := NewSIPResponse(resp)
		if val,ok := parsed.Headers["Call-ID"]; ok && (err == nil) {
			if val == call_id {
				return parsed,nil
			}
		}
	}
}

// Create a TCP connection using source port 5060.
// If we can't nab it, use a random source port.

func ConnectTCP(target_host string, target_port int) (*net.TCPConn,error) {
	target := fmt.Sprintf("%s:%d", target_host, target_port)
	target_addr,err := net.ResolveTCPAddr("tcp4", target)
	if err != nil {
		return nil,err
	}
	local_addr,err := net.ResolveTCPAddr("tcp", ":5060")
	if err != nil {
		return nil,err
	}
	conn,err := net.DialTCP("tcp4", local_addr, target_addr)
	if err == nil {
		return conn,err
	}
	conn,err = net.DialTCP("tcp4", nil, target_addr)
	return conn,err
}

// Send a SIPRequest along an open TCP connection and return the response (or an error).
// You should use a goroutine or conn.SetDeadLine() to ensure it doesn't block forever.

func SendTCP(conn *net.TCPConn, req SIPRequest) error {
	msg := req.Generate()
	_,err := fmt.Fprintf(conn, msg)
	return err
}

func RecvTCP(conn *net.TCPConn) (string,error) {
	output := ""
	for {
		buf := make([]byte, 8192)
		n,err := conn.Read(buf)
		if err != nil {
			return output,err
		}
		if n == 8192 {
			output += string(buf)
		} else {
			output += string(buf[0:n])
			break
		}
	}
	return output,nil
}


