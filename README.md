# Chaos Server

Simple app used in chaos testing as a potential target. It manages value of a counter and allows to access it via REST API.

- [Chaos Server](#chaos-server)
  - [REST API](#rest-api)
  - [Development](#development)

## REST API

- /counter
  - GET: returns counter value
  - POST: increments counter value by one

## Development

To build project:

```shell
go build ./...
```
