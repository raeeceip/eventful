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
    docker run -p 80:80 -e DB_HOST=<mysql-host> -e DB_USER=<mysql-username> -e DB_PASSWORD=<mysql-password> -e DB_PORT=<mysql-port> -e DB_NAME=<mysql-database> -e JWT_SECRET=<your-jwt-secret> eventful-backend
    ```

    Replace `<mysql-host>`, `<mysql-username>`, `<mysql-password>`, `<mysql-port>`, `<mysql-database>`, and `<your-jwt-secret>` with your MySQL and JWT secret details.

5. The server should now be running on `http://localhost`.

## Using Docker Compose

For more complex setups, you can use Docker Compose:

1. Create a `docker-compose.yml` file in the root directory with the following content:

    ```yaml
    version: '3.8'

    services:
      db:
        image: mysql:8.0
        environment:
          MYSQL_ROOT_PASSWORD: rootpassword
          MYSQL_DATABASE: eventful
          MYSQL_USER: user
          MYSQL_PASSWORD: password
        ports:
          - "3306:3306"
        volumes:
          - db_data:/var/lib/mysql

      app:
        build: .
        depends_on:
          - db
        environment:
          DB_HOST: db
          DB_USER: user
          DB_PASSWORD: password
          DB_NAME: eventful
          JWT_SECRET: your_secret_key
        env_file:
          - .env

      nginx:
        image: nginx:latest
        volumes:
          - ./nginx.conf:/etc/nginx/nginx.conf
        ports:
          - "80:80"
        depends_on:
          - app

    volumes:
      db_data:
    ```

2. Build and run the containers with Docker Compose:

    ```shell
    docker-compose up --build
    ```

3. The server should now be running on `http://localhost`.

## Directory Structure

The project directory is organized as follows:

```shell
eventful-backend/
├── auth/
│   └── auth.go                # Contains the authentication middleware using JWT
├── config/
│   └── config.go              # Contains the database initialization and configuration settings
├── handlers/
│   ├── event_handler.go       # Handles HTTP requests related to events
│   ├── role_handler.go        # Handles HTTP requests related to roles
│   ├── team_handler.go        # Handles HTTP requests related to teams
│   ├── user_handler.go        # Handles HTTP requests related to users
│   ├── login_handler.go       # Handles HTTP requests related to logins
│   └── route_handlers.go      # Handles HTTP requests related to routes
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
├── docker-compose.yml         # Docker Compose file to build and run the containers
├── nginx.conf                 # Nginx configuration file for reverse proxy
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
