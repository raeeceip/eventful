# Eventful Application

This project consists of a Go backend and an Electron frontend.

## Backend Setup

1. Install Docker and Docker Compose.
2. Navigate to the `backend` directory.
3. Run: `docker-compose up --build`

The backend will be available at http://localhost:8080

## Frontend Setup

1. Install Node.js and npm.
2. Navigate to the `frontend` directory.
3. Run: `npm install`
4. To start the Electron app, run: `npm start`

## Building the Electron App

To build the Electron app for distribution:

1. Navigate to the `frontend` directory.
2. Run: `npm run build`

This will create distributable packages in the `dist` directory.

## Development

For development, you can run the backend using Docker and the frontend using `npm start`. Make sure to update the `API_URL` in `renderer.js` if your backend is running on a different address.

## API Documentation

For detailed information on the available API endpoints, please refer to the [API documentation](API.md).

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
