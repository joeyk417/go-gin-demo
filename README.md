## 1. Introduction

This repo contains a Web application with a single endpoint: http://localhost:8080/api/equityPositions

In order for the API request to work, you'll need to pass through the following user token as Authorization Bearer Token in header with value: `fJCoxhq8uR9GiUIgaIGfMgw7zCqxwDhQ`.

## 2. Run the API

`go run ./cmd/service`

## 2. Run the Unit Tests

`go test -v ./...`
