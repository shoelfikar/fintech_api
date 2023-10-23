# Backend API Setup and running

This repository contains the Docker Compose configuration to easily run and manage the MyApp application using Docker containers.

## Prerequisites

Before you begin, make sure you have the following installed on your system:

- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Docker Compose: [Installation Guide](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/shoelfikar/fintech_api
   cd backend_api

2. Build Docker image

   ```bash
   docker build -t backend_api:latest .

3. Running Docker Compose
    ```bash
    docker compose up -d

## Manual Start Project
1. Download Dependency

   ```bash
   cd fintech_api
   go mod tidy
   go mod download;go mod vendor;go mod verify

2. Add Database Credential in Env File

   ```bash
   DB_URL=user:password@(ip:port)/database
   MIGRATE_URL=mysql://user:password@tcp(ip:port)/database

3. Running Project
    ```bash
    go run main.go
