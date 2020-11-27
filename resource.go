package main

// Header represents a single header in a resource
type Header struct {
	Key   string
	Value string
}

// Resource represents a resource loaded from disk
type Resource struct {
	Status  int
	Headers []Header
	Body    []byte
}

func NewResource(base, path string) (Resource, error) {
	return Resource{
		Status: 202,
		Headers: []Header{
			{Key: "X-Test", Value: "Testing"},
		},
		Body: []byte("hello"),
	}, nil
}
