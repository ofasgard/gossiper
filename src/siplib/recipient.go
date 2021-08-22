package siplib

// Struct used to keep track of a SIP recipient. Used for both sender and receiver.

type SIPRecipient struct {
	name string
	extension string
	hostname string
	port int
}

func NewSIPRecipient(name string, extension string, hostname string, port int) SIPRecipient {
	rec := SIPRecipient{}
	rec.name = name
	rec.extension = extension
	rec.hostname = hostname
	rec.port = port
	return rec
}


func (sr SIPRecipient) GetURI() string {
	return GenerateURI(sr.hostname, sr.extension)
}
