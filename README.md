# fbs-golang-test-task

A example service that provides a slice of Fibonacci numbers via gRPC and HTTP.

## Installation

It also requires Redis to be installed.
```
go get https://github.com/RinatNamazov/fbs-golang-test-task
cd cmd/app
go build
./app --config ./../../configs/config.yaml
```

You can quickly check the functionality with a query in your browser: `http://localhost:8080/getFibonacciSequence?from=12&to=27`

## Config

Default ports:
* HTTP: `8080`
* gRPC: `8081`

You can also change the credentials for connecting to Redis in the config.

## License

The source code is published under GPLv3, the license is available [here](LICENSE).
