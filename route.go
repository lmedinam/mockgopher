package mockgopher

type Route struct {
	Request  *Request
	Response *Response
}

type Request struct {
	Path    string
	Method  string
	Headers []*Header
}

type Response struct {
	Template string
	Status   int
	Headers  []*Header
}

type Header struct {
	Key   string
	Value string
}
