# ping-pong

This is a simple Go application built using the Fiber web framework. It provides an API endpoint for transforming and returning text in uppercase.

## Table of Contents

- [Roadmap](#roadmap)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)


## Roadmap

For the future development plans and roadmap of this project, please refer to the [Ping Pong Go Roadmap](https://github.com/log101/YARP/blob/main/projects/tr/ping-pong-go.md) document.
## Installation

1. Make sure you have Go installed on your system.
2. Clone this repository: `git clone https://github.com/gnsamine/ping-pong.git`
3. Change to the project directory: `cd ping-pong`
4. Install required dependencies: `go get github.com/gofiber/fiber/v2`

## Usage

1. Run the application: `go run main.go`
2. Access the application in your browser or using tools like `curl`.
3. Use the provided API endpoints to interact with the application.

## API Endpoints

### 1. Get Uppercase JSON

**URL**: `/amine`

**Method**: GET

**Query Parameters**:
- `text` (string): The text to be transformed to uppercase.

**Response**:
- If a valid `text` parameter is provided, the API will return a JSON response containing the transformed text in uppercase.
- If no `text` parameter is provided, a 400 Bad Request status will be returned.

### 2. Get Uppercase String

**URL**: `/`

**Method**: GET

**Query Parameters**:
- `text` (string): The tet-xt to be transformed to uppercase.

**Response**:
- If a valid `text` parameter is provided, the API will return the transformed text in uppercase as a plain text response.
- If no `text` parameter is provided, an empty response with a 400 Bad Request status will be returned.



