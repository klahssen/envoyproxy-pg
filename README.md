# envoyproxy-pg
simple playground to try envoyproxy

## Dependencies
Docker-compose is required.
It will start 3 containers:
1. Envoy proxy (routing traffic to service1 and service2)
2. Two independant containers (service1 and service2) from the same golang http server

## Process

1. Run `make build`
2. Run `make up` to start the proxy and 2 services with docker-compose
3. the proxy (envoy) listens on port 8000 and redirects traffic

## Routing

/service1/<something> will be routed to service 1
/service2/<something> will be routed to service 2
