
# OHLC with Volume & Value

Technical Challenge Part 1.
Created 5 containers using docker
- redis - for datastore
- zookeeper - to manage kafka
- kafka - for messaging queue
- go-grpc - to server grpc server
- go-consumer - to consume message from kafka (producer)


## Pre-Requisites

- Internet Connection to download docker images
- docker
- GRPC client

## Run Locally

- go to part1 directory which contain docker-compose.yml file then run

```bash
  docker compose up
```

- wait until all containers running
- connect grpc with grpc client to ```localhost:8080```
- hit ```InitData``` and wait for processing
- hit ```GetSummary``` to get OHLC with Volume & Value with example request
```
{
  "stockName": "BBCA"
}
```


## Tools Used

- protoc (to generate proto file)
- mockery (for unit test)
- golangci-lint
## Code Architecture

- Config that can be injected to the app using yaml file
- Resource that hold all the necessary depedency connection like datastore (redis, db, etc), mq (kafka, nsq, etc).
- Handler layer to serialization and deserialization
- Usecase layer to handle business logic
- Datalogic layer to handle data logic
- Repository wrapper to call other depedency or storage
## Flow Process

- After all containers running in docker, client can hit grpc `InitData` to load all the json from subsetdata subfolder
- Then app will scan all the json and produce message per row transaction to be consume by consumer
- In consumer, messages will be process to get the summary of OHLC with Volume and Value and store that summary to redis using `HSET` with identifier Stock Name
- Then client can hit `GetSummary` with request of Stock Name to get the latest summary of all the transaction from redis

*note that you need to hit `InitData` and wait for processing first before `GetSummary` otherwise that will be no data return
*you can check some screenshot in `/screenshot` folder and some diagram too