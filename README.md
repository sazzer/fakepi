# FakePI - Fake API Server

FakePI allows a mock API Server of any kind to be written by simply representing all responses as static files on disk.

## Installation

Installation currently requires Go to be installed. The application can then be installed by exeuting:

```
$ go get -u github.com/sazzer/fakepi
```

This will then download, build and install the executable into your GOPATH. Assuming this is correctly on your system path then the executable is now ready to use.

## Usage

FakePI allows you to write a mock HTTP server by providing a set of static files that represent the responses.

This is different to other static servers in that the static files allow for HTTP status codes, headers and payloads to all be included and returned. This means it can be used to represent an API instead of just serving up files.

To execute, you need to provide:

- A port to listen on - defaults to "8000"
- A directory containing the resources to serve - defaults to the current directory

## Resource Structure

Resources are represented in files with a particular structure:

```
200
Content-Type: application/json

{
  "hello": "world"
}
```

The first line is the HTTP status code. Any valid status code can be used here.

Subsequent lines, up to the first blank line, are HTTP headers. Each header must be on a single line, and must follow the HTTP specification of being a key and value separated by a colon.

All lines after the first blank line are then treated as the body of the response.

## Locating Resource Files

By default, the file to load is simply the path from the incoming request. So, for example, a request to `http://localhost:8000/api/users` will attempt to load the resource found in `/api/users`, relative to our base directory.

If this fails then it will attempt to find an "index document". In this case, the above URL will attempt to load `/api/users/_index` instead. This allows for URLs to point to directories and still be able to serve up responses.

## Failure to find a resource

If no resource can be found, the server will return an HTTP 404 Not Found with no headers and no body.
