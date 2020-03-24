[![pipeline status](https://gitlab.com/ncent/auth/badges/master/pipeline.svg)](https://gitlab.com/ncent/auth/commits/master) [![coverage report](https://gitlab.com/ncent/auth/badges/master/coverage.svg)](https://gitlab.com/ncent/auth/commits/master)

# Auth

This module will be responsible for all the Authentication and Authorization for accessing our backend nCent API

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

```
GO
Node
npm
Serverless Framework
go dep
ginkgo
```

### Installing

Install Serverless Framework

```
npm install -g serverless
```

Install node dependencies

```
npm install
```

## Building

```
make build
```

## Running the tests

```
go test ./...
```

or

```
ginkgo ./...
```

## Deployment

Development environment
```
SLS_DEBUG=* serverless deploy --verbose --stage development
```

Production environment
```
SLS_DEBUG=* serverless deploy --verbose --stage production
```

## Built With

* [GO](https://golang.org) - The Language
* [Serverless Framework](https://serverless.com) - Deployment Framework
* [ElasticCache](https://aws.amazon.com/elasticache/) - Used to store data


## Authors

* **Rodrigo Serviuc Pavezi** - *Initial work* - [rodrigopavezi](https://gitlab.com/rodrigopavezi)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details