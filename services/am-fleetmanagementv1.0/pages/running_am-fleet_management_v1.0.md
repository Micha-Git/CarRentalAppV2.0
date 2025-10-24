Running AM-FleetManagementV1.0

1. Create a copy of the `src/.env.example` file and rename it to `src/.env`. Afterward, update the variable settings within this new file.
  
2. In your terminal, change to the `src/` directory and execute the following command: 
    ```
    go run .
    ```
    
# Using Fleet and Customer Service API

This document provides detailed information on the gRPC APIs for the Management of Fleets services within the `fleetmanagement` package.

---


## Installing grpcurl

To install `grpcurl` using Go, make sure you have Go installed on your machine. Then, run the following command:

```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7
```

## Fleet Service

### `AddCarToFleet`

#### Request Format

```protobuf
AddCarToFleetRequest {
  Vin vin = 1;
  string fleetId = 2;
  string location = 3;
}
```

#### Command Line Request

```bash
grpcurl -d="{\"vin\":\"JH4DB1561NS000569\",\"fleetId\":\"1\,\"location\":\"Karlsruhe\"}" -plaintext localhost:50051 fleetmanagement.FleetService/AddCarToFleet
```

#### Response Body

```json
{
    "car": {
        "vin": "JH4DB1561NS000569",
        "brand": "VW",
        "model": "Golf"
    },
    "error": null
}
```
##### Error Response Body

```json
{
    "car": null,
    "error": {
        "message": "Internal",
        "details": "car with VIN JH4DB1561NS000569 does not exist."
    }
}