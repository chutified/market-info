# Market Info
The Market Info is a REST API that provides up-to-date information about the market. More precisely, it is the data of world commodities, currencies and cryptocurrencies. The service allows access to data in JSON via HTTP.

The web app connects 3 microservices and usus gRPC to communicate with all of them, so the service works properly only if other three microservices are also running. To configure the targets (hosts and ports), set the variables in the config.yml file.

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

### Cryptocureency service
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
