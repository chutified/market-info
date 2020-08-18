# Market Info
The Market Info is a REST API that provides up-to-date information about the market. More precisely, it is the data of world commodities, currencies and cryptocurrencies. The service allows access to data in JSON via HTTP.

The web app connects 3 microservices and usus gRPC to communicate with all of them, so the service works properly only if other three microservices are also running. To configure the targets (hosts and ports), set the variables in the config.yml file.

## Dependencies
The REST API server will start without any of these dependencies, but the handler endpoints that use these services will return an error message on requests.

### Commodity service
    - <a href="https://github.com/chutified/commodity-prices" taget="_blank">Source code</a>
    - <a href="https://github.com/chutified/commodity-prices/blob/master/README.md" taget="_blank">Documentation</a>
    - <a href="https://github.com/chutified/commodity-prices/blob/master/Dockerfile" taget="_blank">Dockerfile</a>

### Currency service
    - <a href="https://github.com/chutified/currencies" taget="_blank">Source code</a>
    - <a href="https://github.com/chutified/currencies/blob/master/README.md" taget="_blank">Documentation</a>
    - <a href="https://github.com/chutified/currencies/blob/master/Dockerfile" taget="_blank">Dockerfile</a>

### Cryptocureency service
    - <a href="https://github.com/chutified/crypto-currencies" taget="_blank">Source code</a>
    - <a href="https://github.com/chutified/crypto-currencies/blob/master/README.md" taget="_blank">Documentation</a>
    - <a href="https://github.com/chutified/crypto-currencies" taget="_blank">Dockerfile</a>
