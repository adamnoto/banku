# Banku

A demonstration for event sourcing using Go and Kafka.

To run this project:

1. Install Go
2. Install Kafka
3. Install govendor
4. Run `govendor sync`
5. Run `go build && ./banku --act=consumer` to run the program as consumer
6. Run `go build && ./banku` to run the program as producer

To run testing:

    docker-compose run app ginkgo

Or, you may as well follow the tutorial here:

Cheers, Adam!
