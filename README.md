# Eventful Backend Server

This is the backend server for the Eventful event management system. It is built with Gin and can be started using Docker. The server also uses a MySQL database.

## Prerequisites

Before running the server, make sure you have the following installed:

- Docker
- MySQL

## Getting Started

1. Clone the repository:

    ```shell
    git clone https://github.com/raeeceip/eventful-backend.git
    ```

2. Navigate to the project directory:

    ```shell
    cd eventful-backend
    ```

3. Build the Docker image:

    ```shell
    docker build -t eventful-backend .
    ```

4. Start the Docker container:

    ```shell
    docker run -p 8080:8080 -e DB_HOST=<mysql-host> -e DB_USER=<mysql-username> -e DB_PASSWORD=<mysql-password> -e DB_PORT=<mysql-port> -e DB_NAME=<mysql-database> -e OKTA_CLIENT_ID=<okta-client-id> -e OKTA_ISSUER=<okta-issuer> eventful-backend
    ```

    Replace `<mysql-host>`, `<mysql-username>`, `<mysql-password>`, `<mysql-port>`, `<mysql-database>`, `<okta-client-id>`, and `<okta-issuer>` with your MySQL and Okta details.

5. The server should now be running on `http://localhost:8080`.

## Directory Structure

The project directory is organized as follows:

``shell 
eventful-backend/
├── auth/
│   └── auth.go                # Contains the authentication middleware using Okta
├── config/
│   └── config.go              # Contains the database initialization and configuration settings
├── handlers/
│   ├── event_handler.go       # Handles HTTP requests related to events
│   ├── role_handler.go        # Handles HTTP requests related to roles
│   ├── team_handler.go        # Handles HTTP requests related to teams
│   └── user_handler.go        # Handles HTTP requests related to users
├── models/
│   ├── event.go               # Defines the Event model
│   ├── role.go                # Defines the Role model
│   ├── team.go                # Defines the Team model
│   └── user.go                # Defines the User model
├── repositories/
│   ├── event_repository.go    # Contains database operations for events
│   ├── role_repository.go     # Contains database operations for roles
│   ├── team_repository.go     # Contains database operations for teams
│   └── user_repository.go     # Contains database operations for users
├── .env                       # Contains environment variables (not included in the repository)
├── Dockerfile                 # Dockerfile to build the Docker image
├── go.mod                     # Go module dependencies
├── go.sum                     # Go module dependency checksums
└── main.go                    # Entry point of the application

``` 

## API Documentation

For detailed information on the available API endpoints, please refer to the [API documentation](api.md).

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
