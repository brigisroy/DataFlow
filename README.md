# DataFlow


##Overview

This document provides instructions on how to set up and run the project using the Makefile and Docker Compose configuration provided. The project involves database migrations, running a Go application, and syncing data from cloud storage systems with CSV files in the data folder.

### Prerequisites

Ensure you have the following installed on your system:

- Docker
- Docker Compose 
- Go (Golang)
- goose (for database migrations)

### Steps to Run the Project

1. Clone the Repository

### Clone the repository
> git clone https://github.com/brigisroy/DataFlow.git

### Navigate into the project directory
> cd DataFlow

### Set Up the Database

Using Docker Compose, set up the PostgreSQL database:

### Run the PostgreSQL service using Docker Compose
> docker-compose up -d postgres

This command will start the PostgreSQL service in detached mode.

### Apply Database Migrations

Use the Makefile to apply database migrations using goose:

### Ensure goose is installed and apply migrations
> make migrate

This command will check if goose is installed, install it if necessary, and then apply the migrations located in the db/postgres/schema/ directory.


add csv to the data folder and update the .env for csv file you need to sync 

### Run the Go Application

Run the main.go file using the Makefile:

### Run the Go application
> make run

This command will execute main.go using go run.

Verify Application and Database

Once the application is running, you can verify that it is properly connected to the database and functioning as expected.

Additional Commands

> Migrate Down

To roll back the last migration, use:

> make migrate-down

Check Migration Status

To check the status of migrations, use:

> make status

Stopping Services

To stop the services started by Docker Compose:

> docker-compose down

This command will stop and remove all containers defined in the docker-compose.yml file.

Notes

Ensure that the environment variables in the Makefile and docker-compose.yml are correctly set to match your development setup.

The goose command should be available in your PATH. If it isn't, ensure you install it or add the correct path to the Makefile.

Troubleshooting

goose not recognized: Ensure you have Go installed and that your ``GOPATH/bin`` is included in your ```$PATH```.

Docker services not starting: Check that Docker is running and that there are no port conflicts.

By following these steps, you should be able to successfully run the project, manage the database migrations, and sync data from cloud storage. This assumes the following: periodic database syncing will be performed every 30 minutes, and the CSV files will be present in the data folder. 