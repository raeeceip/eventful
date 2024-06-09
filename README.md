# Eventful Backend Server

This is the backend server for the Eventful event management system. It is built with Gorilla and can be started using Docker. The server also uses a MySQL database.

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
    docker run -p 8080:8080 -e DB_HOST=<mysql-host> -e DB_USER=<mysql-username> -e DB_PASSWORD=<mysql-password> eventful-backend
    ```

    Replace `<mysql-host>`, `<mysql-username>`, and `<mysql-password>` with your MySQL database details.

5. The server should now be running on `http://localhost:8080`.

## API Documentation

For detailed information on the available API endpoints, please refer to the [API documentation](api.md).

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
