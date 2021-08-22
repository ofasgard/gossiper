package siplib

import "strings"
import "strconv"
import "errors"

// Struct used to keep track of a SIP response.

type SIPResponse struct {
	Status string
	StatusCode int
	Headers map[string]string
	Body string
}

func (r *SIPResponse) SetHeader(header string, value string) {
	r.Headers[header] = value
}


func NewSIPResponse(raw_response string) (SIPResponse,error) {
	output := SIPResponse{}
	output.Headers = make(map[string]string)
	//Decapitate the response (i.e. remove head from body).
	parts := strings.Split(raw_response, "\r\n\r\n")
	head := parts[0]
	body := ""
	if (len(parts) > 1) {
		body = parts[1]
	}
	output.Body = body
	//Split the head up into status line and headers.
	headers := strings.Split(head, "\r\n")
	if (len(headers) < 2) {
		return output,errors.New("Failed to parse SIP response: less than 2 header lines.")
	}
	output.Status = headers[0]
	//Attempt to extract SIP response code.
	status_parts := strings.Split(output.Status, " ")
	if (len(status_parts) < 2) {
		return output,errors.New("Failed to parse SIP response: status line makes no sense")
	}
	code,err := strconv.Atoi(status_parts[1])
	if (err != nil) {
		return output,err
	}
	output.StatusCode = code
	//Now it's time to parse the headers.
	for _,header := range headers[1:] {
		header_parts := strings.Split(header, ": ")
		if len(header_parts) > 1 {
			output.SetHeader(header_parts[0], header_parts[1])
		}
	}
	return output,nil
}
