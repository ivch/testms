# Test task for Acme Corp

## Notes
#### Client `parser`
Since data source can be not only file i decided to make separate package `parser` which
send data to shared channel.

I intentionally ommited full struct description in order to win some time :)


#### Server
I've decided to not use any orm in order to try to use less non stdlib packages

## Requirements
- Docker 17.05+
- docker-compose

## Build
- `make all`

## Run
- `docker-compose up`

## Clean
- `make clean`

## Try
`curl -X GET http://localhost:8080/v1/port/AEAUH`