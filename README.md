# Market Info
The Market Info is a REST API that provides up-to-date information about the market. More precisely, it is the data of world commodities, currencies and cryptocurrencies. The service allows access to data in JSON via HTTP.

The web app connects 3 microservices and usus gRPC to communicate with all of them, so the service works properly only if other three microservices are also running. To configure the targets (hosts and ports), set the variables in the config.yml file.
