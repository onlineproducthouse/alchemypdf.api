# AlcheMyPDF API

## Running API

To start the API, first you need to install dependencies by executing following commands:

```bash
 go mod download
 go get -u github.com/swaggo/swag/cmd/swag
 go install github.com/swaggo/swag/cmd/swag
 $(go env GOPATH)/bin/swag init --parseDependency -g services/alchemypdf.api.app/app/app.go
```

Now you can either run the VSCode debugger to execute the project or execute the command:

```bash
# API
$ export $(echo $(cat .env | sed 's/#.*//g'| xargs) | envsubst) && go run ./services/alchemypdf.api.app/main.go
```

In the logs, the API will output something like:

```bash
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [host]:[port]
```

This indicates the API has started.

You can open Swagger with this URL: `http://127.0.0.1:7890/swagger/`

To have environment dependencies set up, execute command:

```bash
$ export $(echo $(cat .env | sed 's/#.*//g'| xargs) | envsubst) &&  docker compose -f 'docker-compose.debug.yml' up --build
```

## Unit tests

There are only two areas focused on for Unit Tests, being the utilities and domain

For testing services:
```bash
$ go test ./lib/alchemypdf.api.service/... -coverprofile=coverage-service.out
$ go tool cover -html=coverage-service.out
```
