# Market Info
The Market Info is a REST API that provides up-to-date information about the market. More precisely, it is the data of world commodities, currencies and cryptocurrencies. The service allows access to data in JSON via HTTP.

The web app connects 3 microservices and usus gRPC to communicate with all of them, so the service works properly only if other three microservices are also running. To configure the targets (hosts and ports), set the variables in the config.yml file.

## Table of content
    - Dependencies
        - Commodity service
        - Currency service
        - Cryptocurency service
    - Tools
    - API sources
    - API documentation
    - Usage
    - Examples
    - Configuration
    - Directory structure

## Dependencies
The REST API server will start without any of these dependencies, but the handler endpoints that use these services will return an error message on requests.

### Commodity service
    - <a href="https://github.com/chutified/commodity-prices" taget="_blank">Source code</a>
    - <a href="https://github.com/chutified/commodity-prices/blob/master/README.md" taget="_blank">Documentation</a>
    - <a href="https://github.com/chutified/commodity-prices/blob/master/Dockerfile" taget="_blank">Dockerfile</a>
    - <a href="https://github.com/chutified/commodity-prices#supported-commodities" target="_blank">Supported commodities</a>

### Currency service
    - <a href="https://github.com/chutified/currencies" taget="_blank">Source code</a>
    - <a href="https://github.com/chutified/currencies/blob/master/README.md" taget="_blank">Documentation</a>
    - <a href="https://github.com/chutified/currencies/blob/master/Dockerfile" taget="_blank">Dockerfile</a>
    - <a href="https://github.com/chutified/currencies/blob/master/README.md#supported-currency-codes" target="_blank">Supported currencies</a>

### Cryptocurency service
    - <a href="https://github.com/chutified/crypto-currencies" taget="_blank">Source code</a>
    - <a href="https://github.com/chutified/crypto-currencies/blob/master/README.md" taget="_blank">Documentation</a>
    - <a href="https://github.com/chutified/crypto-currencies" taget="_blank">Dockerfile</a>
    - <a href="https://github.com/chutified/crypto-currencies/blob/master/docs/currencies.md" target="_blank">Supported cryptocurrencies</a>

## Tools
    - <a href="https://gin-gonic.com/" target="_blank">Gin framework</a>
    - <a href="https://grpc.io/" target="_blank">gRPC</a>
    - <a href="https://git-scm.com/" target="_blank">Git</a>
    - <a href="https://www.docker.com/" target="_blank">Docker Engine</a>
    - <a href="https://swagger.io/" target="_blank">Swagger 2.0</a>

## API sources
    - Commodity (<a href="https://markets.businessinsider.com/currencies" target="_blank">Markets - Business Insider</a>)
    - Currency (<a href="https://markets.businessinsider.com/currencies" target="_blank">Markets - Business Insider</a>)
    - Cryptocurency (<a href="https://coinmarketcap.com/all/views/all/" target="_blank">CoinMarketCap</a>)

## API documentation
The API uses the industry standart Swagger tool for its documentation.
Swagger documentation file: <a href="https://github.com/chutified/market-info/blob/master/docs/swagger.json" target="_blank">JSON</a>/<a href="https://github.com/chutified/market-info/blob/master/docs/swagger.yaml" target="_blank">YAML</a>

## Usage

## Examples

## Configuration
**Default:**
```yaml
---
api_port: 80
commodity_service_target: 'localhost:10501'
currency_service_target: 'localhost:10502'
crypto_service_target: 'localhost:10503'
```
*api_port:* the port of the host on which the server is running
*commodity_service_target:* the "host:post" of the commodity microservice server
*currency_service_target:* the "host:post" of the currency microservice server
*crypto_service_target:* the "host:post" of the cryptocurrency microservice server

## Directory structure
```bash
_
├── config
│   └── config.go
├── data
│   ├── commodity_service.go
│   ├── crypto_service.go
│   └── currency_service.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── handlers
│   ├── commodities.go
│   ├── cryptos.go
│   ├── currencies.go
│   ├── errors.go
│   └── handler.go
├── models
│   ├── commodity.go
│   ├── crypto.go
│   └── currency.go
├── server
│   ├── routes.go
│   └── server.go
├── config.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
├── README.md
└── server.log
```
