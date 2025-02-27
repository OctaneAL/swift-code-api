# swift-code-api

## Description

A SWIFT code, also known as a Bank Identifier Code (BIC), is a unique identifier of a bank's branch or headquarter. It ensures that international wire transfers are directed to the correct bank and branch, acting as a bank's unique address within the global financial network.

This application transforms SWIFT code data stored in spreadsheets into a structured, accessible format for integration with other applications. It enables efficient retrieval of bank information, ensuring seamless international transactions by providing a reliable source for SWIFT/BIC code lookups.

## Install

  ```
  git clone github.com/OctaneAL/swift-code-api
  cd swift-code-api
  go build main.go
  export KV_VIPER_FILE=./config.yaml
  ./main migrate up
  ./main run service
  ```

## Running from docker 
  
Make sure that docker installed.

use `docker run ` with `-p 8080:80` to expose port 80 to 8080

  ```
  docker build -t github.com/OctaneAL/swift-code-api .
  docker run -e KV_VIPER_FILE=/config.yaml github.com/OctaneAL/swift-code-api
  ```

## Running from Source

* Set up environment value with config file path `KV_VIPER_FILE=./config.yaml`
* Provide valid config file
* Launch the service with `migrate up` command to create database schema
* Launch the service with `run service` command


### Database
For services, we do use ***PostgresSQL*** database. 
You can [install it locally](https://www.postgresql.org/download/) or use [docker image](https://hub.docker.com/_/postgres/).


### Documentation

API Endpoints:
...
