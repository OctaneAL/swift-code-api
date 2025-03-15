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


## Documentation

### Overview
The **Swift Code API** allows users to retrieve, add, and delete SWIFT codes, as well as fetch SWIFT codes by country.

### Base URL
```
http://localhost:8000/v1/swift-codes
```
### Get SWIFT Code Details

**Endpoint:**  
```http
GET /v1/swift-codes/{swiftCode}
```
Retrieves details of a SWIFT code, including bank name, country, and branches if applicable.

#### Successful Response (200 OK)  
If the SWIFT code exists, the response contains details about the bank and its branches (if the SWIFT code belongs to the headquarter).  

**Example Request:**  
```http
GET /v1/swift-codes/BCHICLRMXXX
```

**Example Response:**
```json
{
    "address": "RUE GRIMALDI  MONACO, MONACO, 98000",
    "bankName": "BANCO DE CHILE",
    "countryISO2": "CL",
    "countryName": "CHILE",
    "isHeadquarter": true,
    "swiftCode": "BCHICLRMXXX",
    "branches": [
        {
            "address": "JATKOWA 3  GROJEC, MAZOWIECKIE, 05-600",
            "bankName": "BANCO DE CHILE",
            "countryISO2": "CL",
            "countryName": "CHILE",
            "isHeadquarter": false,
            "swiftCode": "BCHICLRMEXP"
        },
        {
            "address": "SREBARNA 16 PARK LANE OFFICE CENTER, FLOOR 6 SOFIA, SOFIA, 1407",
            "bankName": "BANCO DE CHILE",
            "countryISO2": "CL",
            "countryName": "CHILE",
            "isHeadquarter": false,
            "swiftCode": "BCHICLRMIMP"
        }
    ]
}
```
If the requested SWIFT code is not a headquarter, the `branches` array will be empty.

#### Error Response (404 Not Found)  
If the SWIFT code is not found, the API returns an error.

**Example Response:**
```json
{
    "errors": [
        {
            "title": "Swift code not found",
            "status": "404"
        }
    ]
}
```

### Get SWIFT Codes by Country

**Endpoint:**  
```http
GET /v1/swift-codes/country/{countryISO2code}
```
Retrieves all SWIFT codes available for a given country.

#### Successful Response (200 OK)  
If SWIFT codes are available for the specified country, the response contains a list of banks and their SWIFT codes.

**Example Request:**  
```http
GET /v1/swift-codes/country/PL
```

**Example Response:**
```json
{
    "countryISO2": "PL",
    "countryName": "POLAND",
    "swiftCodes": [
        {
            "address": "UL. WRONIA 31  WARSZAWA, MAZOWIECKIE, 00-846",
            "bankName": "BNP PARIBAS S.A. BRANCH IN POLAND",
            "countryISO2": "PL",
            "countryName": "POLAND",
            "isHeadquarter": true,
            "swiftCode": "BNPAPLPXXXX"
        },
        {
            "address": "UL. CYPRIANA KAMILA NORWIDA 1  GDANSK, POMORSKIE, 80-280",
            "bankName": "BANK BPH SA",
            "countryISO2": "PL",
            "countryName": "POLAND",
            "isHeadquarter": false,
            "swiftCode": "BPHKPLPKCUS"
        }
    ]
}
```

#### Error Response (404 Not Found)  
If no SWIFT codes are found for the given country, the API returns an error.

**Example Request:**  
```http
GET /v1/swift-codes/country/UA
```

**Example Response:**
```json
{
    "errors": [
        {
            "title": "Not Found",
            "status": "404"
        }
    ]
}
```

### Get SWIFT Code Details

**Endpoint:**  
```http
GET /v1/swift-codes/{swiftCode}
```
Retrieves details of a SWIFT code, including bank name, country, and branches if applicable.

#### Successful Response (200 OK)  
If the SWIFT code exists, the response contains details about the bank and its branches (if the SWIFT code belongs to the headquarter).  

**Example Request:**  
```http
GET /v1/swift-codes/BCHICLRMXXX
```

**Example Response:**
```json
{
    "address": "RUE GRIMALDI  MONACO, MONACO, 98000",
    "bankName": "BANCO DE CHILE",
    "countryISO2": "CL",
    "countryName": "CHILE",
    "isHeadquarter": true,
    "swiftCode": "BCHICLRMXXX",
    "branches": [
        {
            "address": "JATKOWA 3  GROJEC, MAZOWIECKIE, 05-600",
            "bankName": "BANCO DE CHILE",
            "countryISO2": "CL",
            "countryName": "CHILE",
            "isHeadquarter": false,
            "swiftCode": "BCHICLRMEXP"
        },
        {
            "address": "SREBARNA 16 PARK LANE OFFICE CENTER, FLOOR 6 SOFIA, SOFIA, 1407",
            "bankName": "BANCO DE CHILE",
            "countryISO2": "CL",
            "countryName": "CHILE",
            "isHeadquarter": false,
            "swiftCode": "BCHICLRMIMP"
        }
    ]
}
```
If the requested SWIFT code is not a headquarter, the `branches` array will be empty.

#### Error Response (404 Not Found)  
If the SWIFT code is not found, the API returns an error.

**Example Response:**
```json
{
    "errors": [
        {
            "title": "Swift code not found",
            "status": "404"
        }
    ]
}
```

### Get SWIFT Codes by Country

**Endpoint:**  
```http
GET /v1/swift-codes/country/{countryISO2code}
```
Retrieves all SWIFT codes available for a given country.

#### Successful Response (200 OK)  
If SWIFT codes are available for the specified country, the response contains a list of banks and their SWIFT codes.

**Example Request:**  
```http
GET /v1/swift-codes/country/PL
```

**Example Response:**
```json
{
    "countryISO2": "PL",
    "countryName": "POLAND",
    "swiftCodes": [
        {
            "address": "UL. WRONIA 31  WARSZAWA, MAZOWIECKIE, 00-846",
            "bankName": "BNP PARIBAS S.A. BRANCH IN POLAND",
            "countryISO2": "PL",
            "countryName": "POLAND",
            "isHeadquarter": true,
            "swiftCode": "BNPAPLPXXXX"
        },
        {
            "address": "UL. CYPRIANA KAMILA NORWIDA 1  GDANSK, POMORSKIE, 80-280",
            "bankName": "BANK BPH SA",
            "countryISO2": "PL",
            "countryName": "POLAND",
            "isHeadquarter": false,
            "swiftCode": "BPHKPLPKCUS"
        }
    ]
}
```

#### Error Response (404 Not Found)  
If no SWIFT codes are found for the given country, the API returns an error.

**Example Request:**  
```http
GET /v1/swift-codes/country/UA
```

**Example Response:**
```json
{
    "errors": [
        {
            "title": "Not Found",
            "status": "404"
        }
    ]
}
```

### Add SWIFT Code

**Endpoint:**  
```http
POST /v1/swift-codes/
```
Adds a new SWIFT code to the database.

#### Successful Response (201 Created)  
If the SWIFT code is successfully created, the API returns a confirmation message.

**Example Request:**  
```json
{
    "address": "123 Main Street, London, UK, 10001",
    "bankName": "HSBC BANK PLC",
    "countryISO2": "GB",
    "countryName": "UNITED KINGDOM",
    "isHeadquarter": true,
    "swiftCode": "HBUKGB4BXXX"
}
```

**Example Response:**
```json
{
    "message": "Swift code created successfully"
}
```

#### Error Response (400 Bad Request)  
If the request data is invalid, the API returns an error indicating which field failed validation.

**Example Response:**
```json
{
    "errors": [
        {
            "title": "Bad Request",
            "status": "400",
            "meta": {
                "error": "the length must be exactly 2",
                "field": "countryISO2"
            }
        }
    ]
}
```

### Delete SWIFT Code

**Endpoint:**  
```http
DELETE /v1/swift-codes/{swiftCode}
```
Deletes a SWIFT code from the database.

#### Successful Response (200 OK)  
If the SWIFT code is successfully deleted, the API returns a confirmation message.

**Example Request:**  
```http
DELETE /v1/swift-codes/AABBAAAXXX
```

**Example Response:**
```json
{
    "message": "Swift code deleted successfully"
}
```


