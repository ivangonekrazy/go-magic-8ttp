Magic 8ttp-Ball Service
=======================

Returns [Magic 8-ball](http://en.wikipedia.org/wiki/Magic_8-Ball) fortunes "RESTfully".

Building and Running
--------------------

The [Go compiler](http://golang.org/) is required to build this service.

To build, run the Go build command: `go build go-magic-8ttp.go`.

To start the service: `go-magic-8ttp`

To build and run in one step: `go run go-magic-8ttp.go`

The service will listen for HTTP requests on port 8080 and respond with Magic 8-Ball
fortunes along a corresponding HTTP status code.

Ctrl-C will stop the service.

[Magic 8-ball fortunes cited from Wikipedia](http://en.wikipedia.org/wiki/Magic_8-Ball#Possible_answers).
