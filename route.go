package mockgopher

type Route struct {
	Request  *Request
	Response *Response
}

type Request struct {
	Path    string
	Methods []string
	Headers []*Header
}

type Response struct {
	Template  string
	Status    uint16
	Resources []string
	Delay     *int64
	Headers   []*Header
}

type Header struct {
	Key   string
	Value string
}
