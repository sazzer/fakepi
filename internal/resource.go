package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func NewResource(path string) (*Resource, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open resource %s: %w", path, err)
	}
	defer file.Close()

	result := Resource{}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		// Status code
		line := scanner.Text()
		result.Status, err = strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse status code %s: %w", line, err)
		}
	} else {
		return nil, fmt.Errorf("Resource was empty")
	}

	for scanner.Scan() {
		// Headers
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			result.Headers = append(result.Headers, Header{
				Key:   strings.TrimSpace(parts[0]),
				Value: strings.TrimSpace(parts[1]),
			})
		}
	}

	for scanner.Scan() {
		// Body
		line := scanner.Bytes()
		result.Body = append(result.Body, line...)
		result.Body = append(result.Body, []byte("\n")...)
	}

	return &result, nil
}
