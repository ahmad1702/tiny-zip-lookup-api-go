# Zip Info Service

This service provides information about ZIP codes in the United States. At some point the plan is to add international zip code support.

## Features

- Look up ZIP code information via HTTP API.
- Log incoming requests.

## Information Provided

1. _`Zip` (int)_: The actual ZIP code, represented as an integer. This is the primary identifier for the record.
2. _`City` (string)_: The name of the city associated with the ZIP code. This helps users identify the urban area that the ZIP code covers.
3. _`State` (string)_: The full name of the state where the ZIP code is located. This provides geographical context and is essential for users to understand the location.
4. _`StateAbbr` (string)_: The two-letter abbreviation for the state (e.g., "CA" for California). This is useful for quick reference and is commonly used in forms and applications.
5. _`County` (string)_: The name of the county that the ZIP code falls within. This information can be important for demographic studies, local governance, and regional services.
6. _`CountyCode` (int)_: A numerical code representing the county. This can be used for internal processing or referencing in databases.
7. _`Latitude` (float64)_: The geographical latitude of the center point of the ZIP code area. This is useful for mapping applications and geographic information systems (GIS).
8. _`Longitude` (float64)_: The geographical longitude of the center point of the ZIP code area. Like latitude, this is essential for mapping and location-based services.

## Getting Started

### Prerequisites

- Go 1.23.2 or higher

### Running Locally

Clone the repository and run the following commands:

```bash
go mod tidy
go run main.go
```

## Running with Docker

To build and run the application using Docker, follow these steps:

1. **Build the Docker image**:

   ```bash
   docker build -t zip-info-service .
   ```

2. **Run the Docker container**:

   ```bash
   docker run -p 4321:4321 zip-info-service
   ```

   This command maps port `4321` of the container to port `4321` on your host machine.

## API Endpoints

- `GET /`: Returns "Hello World!"
- `GET /zip-info/{zip}`: Returns information about the specified ZIP code.

## Deployment

This application is containerized and can be deployed on any container service like Railway.

## Contributing

Contributions are welcome. Please feel free to submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
