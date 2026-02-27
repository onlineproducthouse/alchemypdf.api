# AlcheMyPDF API

AlcheMyPDF API project

## Dependencies

A running instance of PostreSQL is required and [database migrations](https://github.com/onlineproducthouse/alchemypdf.db/blob/main/README.md) must have been executed.

## Installation

```bash
# clone repository
mkdir alchemypdf.api
cd alchemypdf.api
git clone https://github.com/onlineproducthouse/alchemypdf.api.git .

# install dependencies
go mod download

# install swagger
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
```

## Usage

```bash
# set to either: local, test, qa, prod
export ENVIRONMENT_NAME=local

# set environment variables for the API
export RUN_SWAGGER=true

export ALCHEMYPDF_API_HOST=127.0.0.1
export ALCHEMYPDF_API_KEYS=69d2eddc-2cc9-acab-1a9c-dfcb1fca3efb
export ALCHEMYPDF_API_PORT=10000
export ALCHEMYPDF_API_PROTOCOL=http
export ALCHEMYPDF_PROJECT_NAME=AlcheMyPDF
export ALCHEMYPDF_PROJECT_SHORT_NAME=AlcheMyPDF

# initialise swagger docs
$(go env GOPATH)/bin/swag init --parseDependency -g services/alchemypdf.api.app/app/app.go

# run API
go run ./services/alchemypdf.api.app/main.go
```

In the logs, the API will output something like:

```bash
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ vX.X.X
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [host]:[port]
```

This indicates the API has started.

You can open Swagger with this URL: `http://127.0.0.1:10000/swagger/`

## Unit tests

There are only two areas focused on for Unit Tests, being the utilities and domain

For testing services:
```bash
go test ./lib/alchemypdf.api.service/... -coverprofile=coverage-service.out
go tool cover -html=coverage-service.out
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
