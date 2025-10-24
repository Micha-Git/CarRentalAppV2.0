## Running AM-RentalManagementV2.0

1. Create a copy of the `src/.env.example` file and rename it to `src/.env`. Afterward, update the variable settings within this new file.
  
2. In your terminal, change to the `src/` directory and execute the following command: 
    ```
    go run .
    ```


# Rental, Customer and RentableCars Service API

## Overview

This document provides detailed information on the gRPC APIs for the Rental and Customer services within the `rentalmanagement` package.

---

## Installing grpcurl

To install `grpcurl` using Go, make sure you have Go installed on your machine. Then, run the following command:

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7
```

## RentalsCollectionService

### `ListAvailableCars`

#### Request Format

```protobuf
ListAvailableCarsRequest {
    startDate: google.protobuf.Timestamp;
    endDate: google.protobuf.Timestamp;
    location: string;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"startDate\":\"2022-02-19T00:00:00Z\",\"endDate\":\"2022-02-20T00:00:00Z\",\"location\":\"Karlsruhe\"}" -plaintext localhost:9001 rentalmanagement.RentalsCollectionService/ListAvailableCars
```

#### Response Body

```json
{
  "cars": [
    {
      "vin": {
        "vin": "JH4DB1561NS000565"
      },
      "brand": "VW",
      "model": "ID2",
      "location": "Karlsruhe",
      "pricePerDay": 20
    }
  ],
  "error": null
}
```

##### Error Response Body

```json
{
  "cars": [],
  "error": {
    "message": "Internal",
    "details": "StartDate must be before EndDate"
  }
}
```


### `ListCarRentals`

#### Request Format

```protobuf
ListCarRentalsRequest {
    vin: Vin;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"vin\":\"{\"startDate\":\"JH4DB1561NS000565\"}\"}" -plaintext localhost:9001 rentalmanagement.RentalsCollectionService/ListCarRentals
```

#### Response Body

```json
{
  "rentals": [
    {
      "id": "efc1edc7-c5e4-4f02-8b8e-29aae4ce2c5c",
      "startDate": "2022-02-19T00:00:00Z",
      "endDate": "2022-02-20T00:00:00Z",
      "car": {
        "vin": {
          "vin": "JH4DB1561NS000565"
        },
        "brand": "VW",
        "model": "ID2",
        "location": "Karlsruhe",
        "pricePerDay": 20
      },
      "price": 20,
      "customerId": "customer_ID"
    }
  ],
  "error": null
}
```

##### Error Response Body

```json
{
  "rentals": [],
  "error": {
    "message": "Internal",
    "details": "VIN can not be empty"
  }
}
```

### `ListCustomerRentals`

#### Request Format

```protobuf
ListCustomerRentalsRequest {
    customerId: string;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"CustomerID\":\"customer_ID\"}" -plaintext localhost:9001 rentalmanagement.RentalsCollectionService/ListCustomerRentals
```

#### Response Body

```json
{
  "rentals": [
    {
      "id": "efc1edc7-c5e4-4f02-8b8e-29aae4ce2c5c",
      "startDate": "2022-02-19T00:00:00Z",
      "endDate": "2022-02-20T00:00:00Z",
      "car": {
        "vin": {
          "vin": "JH4DB1561NS000565"
        },
        "brand": "VW",
        "model": "ID2",
        "location": "Karlsruhe",
        "pricePerDay": 20
      },
      "price": 20,
      "customerId": "customer_ID"
    }
  ],
  "error": null
}
```

##### Error Response Body

```json
{
  "rentals": [],
  "error": {
    "message": "InvalidArgument",
    "details": "Customer ID is invalid : "
  }
}
```

## RentableCarsCollectionService

### `AddCarToRental`

#### Request Format

```protobuf
AddCarToRentalRequest {
    vin: Vin;
    location: string;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"vin\":\"{\"startDate\":\"JH4DB1561NS000565\"}\", \"location\":\"Karlsruhe\"}" -plaintext localhost:9001 rentalmanagement.RentableCarsCollectionService/AddCarToRental
```

#### Response Body

```json
{
  "car": {
    "vin": {
      "vin": "JH4DB1561NS000565"
    },
    "brand": "VW",
    "model": "ID2",
    "location": "Karlsruhe",
    "pricePerDay": 20
  },
  "error": null
}
```

##### Error Response Body

```json
{
  "car": null,
  "error": {
    "message": "InvalidArgument",
    "details": "Vin, or location : vin:{vin:\"JH4DB1561NS000565\"}"
  }
}
```

### `RemoveRentableCar`

#### Request Format

```protobuf
RemoveRentableCarRequest {
    vin: Vin;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"vin\":\"{\"startDate\":\"JH4DB1561NS000565\"}\"}" -plaintext localhost:9001 rentalmanagement.rentalmanagement.RentableCarsCollectionService/RemoveRentableCar
```

#### Response Body

```json
{
  "result": true,
  "error": null
}
```

##### Error Response Body

```json
{
  "result": false,
  "error": {
    "message": "Internal",
    "details": "Failed to remove rentable car with VIN JH4DB1561NS000565: Database failed to delete rentable car with VIN JH4DB1561NS000565: record not found"
  }
}
```

## Customer Service

### `RentCar`

#### Request Format

```protobuf
RentCarRequest {
    customerId: string;
    startDate: google.protobuf.Timestamp;
    endDate: google.protobuf.Timestamp;
    vin: Vin;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"customerId\":\"customer_ID\", \"startDate\":\"2022-02-19T00:00:00Z\",\"endDate\":\"2022-02-20T00:00:00Z\",\"vin\":\"{\"startDate\":\"JH4DB1561NS000565\"}\"}" -plaintext localhost:9001 rentalmanagement.CustomerService/RentCar
```

#### Response Body

```json
{
  "rental": {
    "id": "efc1edc7-c5e4-4f02-8b8e-29aae4ce2c5c",
    "startDate": "2022-02-19T00:00:00Z",
    "endDate": "2022-02-20T00:00:00Z",
    "car": {
      "vin": {
        "vin": "JH4DB1561NS000565"
      },
      "brand": "VW",
      "model": "ID2",
      "location": "Karlsruhe",
      "pricePerDay": 20
    },
    "price": 20,
    "customerId": "JH4DB1561NS000565"
  },
  "error": null
}
```

##### Error Response Body

```json
{
  "rental": null,
  "error": {
    "message": "InvalidArgument",
    "details": "VIN, start date, customerId or end date is not valid : startDate:{seconds:1645228800} endDate:{seconds:1645315200} vin:{vin:\"JH4DB1561NS000565\"}"
  }
}
```

### `CancelRental`

```protobuf
CancelRentalRequest {
    customerId: string;
    rentalId: string;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"customerId\":\"customer_ID\", \"rentalId\":\"efc1edc7-c5e4-4f02-8b8e-29aae4ce2c5c\"}" -plaintext localhost:9001 rentalmanagement.CustomerService/CancelRental
```

#### Response Body

```json
{
  "result": true,
  "error": null
}
```

##### Error Response Body

```json
{
  "result": false,
  "error": {
    "message": "Internal",
    "details": "Rental with ID efc1edc7-c5e4-4f02-8b8e-29aae4ce2c5c does not belong to customer with ID customer_ID_2"
  }
}
```