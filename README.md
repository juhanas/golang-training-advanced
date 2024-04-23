# golang-training-advanced

## Description
This repository contains example code for learning advanced techniques of Golang. The code is meant to be read from branch-to-branch, providing increasing detail towards the final solution. Some changes may not be visible in the final solution, thus it is adviced to not skip ahead if you wish to learn the most.

The final branch in this repository aims to show a mid-level production-ready web server. However, there will be certain items lacking, such as exhaustive tests.

Note: The base code is "broken" on purpose, and several tests fail. These will be fixed in subsequent solution branches.

## Installation
1. Clone the repo
2. Run `go mod tidy` to download and install dependencies

## Usage
1. Start the server with `go run .`
2. Make GET request to `127.0.0.1:8080/` to test the server (returns 404)
    - Use for example `curl http://localhost:8080`

### Endpoints available

- GET /find-word?word=cat
    - Returns the number of times the given word occurs in the book data stored on the server
- POST /secrets
    - Adds a new secret in the database and returns the encrypted value. Expects json-object with "name" and "value" strings.
- GET /secrets/[secretName]
    - Returns the decrypted value of the secret with the given name. Returns error if secret not found.
- GET /secrets/count
    - Returns the number of times all secrets have been read or created.

## Testing
### Unit tests
To run all unit tests: `go test ./...`

To run tests in the current folder: `go test .`

To run tests in another folder: `go test ./path/to/folder`

To run a single test: `go test -run=[TestName] ./...`

To run tests with extra info: `go test -v ./...`