package siplib

import "fmt"

// Struct used to construct a SIP request.

type SIPRequest struct {
	Proto string
	Method string
	URI string
	SIPVersion string
	Headers map[string]string
	Body string
	
	sender SIPRecipient
	receiver SIPRecipient
}

func (r SIPRequest) Generate() string {
	out := ""
	//Generate request line.
	out += fmt.Sprintf("%s %s %s\r\n", r.Method, r.URI, r.SIPVersion)
	//Generate header lines.
	for key, value := range r.Headers {
		out += fmt.Sprintf("%s: %s\n", key, value)
	}
	//The body of a request is optional.
	if (len(r.Body) > 0) {
		out += "\r\n"
		out += r.Body
	}
	return out
}

func NewSIPRequest() SIPRequest {
	req := SIPRequest{}
	req.InitHeaders()
	return req
}

func (r *SIPRequest) InitHeaders() {
	r.Headers = make(map[string]string)
	required_headers := []string{"Via", "To", "From", "Call-ID", "CSeq", "Contact", "Content-Type", "Content-Length"}
	for _,header := range required_headers {
		r.SetHeader(header, "")
	}
}

// Methods used to assist in the generation of SIP requests.

func (r *SIPRequest) SetRequestLine(proto string, method string, host string, extension string) {
	r.Proto = proto
	r.Method = method
	r.URI = GenerateURI(host, extension)
	r.SIPVersion = "SIP/2.0"
}

func (r *SIPRequest) SetHeader(header string, value string) {
	r.Headers[header] = value
}

func (r *SIPRequest) SetBody(body string) {
	r.Body = body
	r.SetHeader("Content-Length", fmt.Sprintf("%d", len(r.Body)))
}

// Functions for generating "template" requests for various methods.

func NewOptionsRequest(proto string, sender SIPRecipient, receiver SIPRecipient) SIPRequest {
	req := NewSIPRequest()
	req.SetRequestLine(proto, "OPTIONS", receiver.hostname, receiver.extension)
	req.SetHeader("Accept", "application/sdp")
	req.SetHeader("Content-Type", "application/sdp")
	req.SetHeader("User-Agent", "Avaya SIP R2.2 Endpoint Brcm Callctrl/1.5.1.0 MxSF/v3.2.6.26:")
	req.SetHeader("To", fmt.Sprintf("%s <%s>", receiver.name, receiver.GetURI()))
	req.SetHeader("From", fmt.Sprintf("%s <%s>;tag=%s", sender.name, sender.GetURI(), random_number_string(46)))
	req.SetHeader("Via", fmt.Sprintf("%s/%s %s:%d;branch=z9hG4bK%s", req.SIPVersion, req.Proto, sender.hostname, sender.port, random_number_string(10)))
	req.SetHeader("Contact", fmt.Sprintf("%s <%s>", sender.name, sender.GetURI()))
	req.SetHeader("CSeq", fmt.Sprintf("1 %s", req.Method))
	req.SetHeader("Call-ID", random_number_string(24))
	req.SetHeader("Max-Forwards", "70")
	req.SetHeader("Content-Length", "0")
	return req
}



