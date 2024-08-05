
### api.md


# API Documentation

This document provides information about the API endpoints available in the Eventful Backend Server.

## Authentication

### Login

**Endpoint:** `POST /login`

**Description:** Authenticates a user and returns a JWT token.

**Request Body:**

```json
{
  "username": "string",
  "password": "string"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "token": "string"
  }
  ```

- **401 Unauthorized**

  ```json
  {
    "error": "Unauthorized"
  }
  ```

## Events

### Create Event

**Endpoint:** `POST /events`

**Description:** Creates a new event.

**Request Body:**

```json
{
  "title": "string",
  "description": "string",
  "date": "YYYY-MM-DD",
  "location": "string"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "title": "string",
    "description": "string",
    "date": "YYYY-MM-DD",
    "location": "string"
  }
  ```

### Get Event by ID

**Endpoint:** `GET /events/:id`

**Description:** Retrieves an event by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "title": "string",
    "description": "string",
    "date": "YYYY-MM-DD",
    "location": "string"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "Event not found"
  }
  ```

### Update Event

**Endpoint:** `PUT /events/:id`

**Description:** Updates an existing event.

**Request Body:**

```json
{
  "title": "string",
  "description": "string",
  "date": "YYYY-MM-DD",
  "location": "string"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "title": "string",
    "description": "string",
    "date": "YYYY-MM-DD",
    "location": "string"
  }
  ```

### Delete Event

**Endpoint:** `DELETE /events/:id`

**Description:** Deletes an event by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "message": "Event deleted"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "Event not found"
  }
  ```

## Roles

### Create Role

**Endpoint:** `POST /roles`

**Description:** Creates a new role.

**Request Body:**

```json
{
  "name": "string",
  "leader": "boolean"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "name": "string",
    "leader": "boolean"
  }
  ```

### Get Role by ID

**Endpoint:** `GET /roles/:id`

**Description:** Retrieves a role by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "name": "string",
    "leader": "boolean"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "Role not found"
  }
  ```

### Get All Roles

**Endpoint:** `GET /roles`

**Description:** Retrieves all roles.

**Response:**

- **200 OK**

  ```json
  [
    {
      "id": "integer",
      "name": "string",
      "leader": "boolean"
    }
  ]
  ```

### Update Role

**Endpoint:** `PUT /roles/:id`

**Description:** Updates an existing role.

**Request Body:**

```json
{
  "name": "string",
  "leader": "boolean"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "name": "string",
    "leader": "boolean"
  }
  ```

### Delete Role

**Endpoint:** `DELETE /roles/:id`

**Description:** Deletes a role by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "message": "Role deleted"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "Role not found"
  }
  ```

## Teams

### Create Team

**Endpoint:** `POST /teams`

**Description:** Creates a new team.

**Request Body:**

```json
{
  "name": "string",
  "description": "string"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "name": "string",
    "description": "string"
  }
  ```

### Get Team by ID

**Endpoint:** `GET /teams/:id`

**Description:** Retrieves a team by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "name": "string",
    "description": "string"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "Team not found"
  }
  ```

### Get All Teams

**Endpoint:** `GET /teams`

**Description:** Retrieves all teams.

**Response:**

- **200 OK**

  ```json
  [
    {
      "id": "integer",
      "name": "string",
      "description": "string"
    }
  ]
  ```

### Update Team

**Endpoint:** `PUT /teams/:id`

**Description:** Updates an existing team.

**Request Body:**

```json
{
  "name": "string",
  "description": "string"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "name": "string",
    "description": "string"
  }
  ```

### Delete Team

**Endpoint:** `DELETE /teams/:id`

**Description:** Deletes a team by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "message": "Team deleted"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "Team not found"
  }
  ```

## Users

### Create User

**Endpoint:** `POST /users`

**Description:** Creates a new user.

**Request Body:**

```json
{
  "username": "string",
  "password": "string",
  "email": "string",
  "role": "string"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "username": "string",
    "email": "string",
    "role": "string"
  }
  ```

### Get User by ID

**Endpoint:** `GET /users/:id`

**Description:** Retrieves a user by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "username": "string",
    "email": "string",
    "role": "string"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "User not found"
  }
  ```

### Get All Users

**Endpoint:** `GET /users`

**Description:** Retrieves all users.

**Response:**

- **200 OK**

  ```json
  [
    {
      "id": "integer",
      "username": "string",
      "email": "string",
      "role": "string"
    }
  ]
  ```

### Update User

**Endpoint:** `PUT /users/:id`

**Description:** Updates an existing user.

**Request Body:**

```json
{
  "username": "string",
  "email": "string",
  "role": "string"
}
```

**Response:**

- **200 OK**

  ```json
  {
    "id": "integer",
    "username": "string",
    "email": "string",
    "role": "string"
  }
  ```

### Delete User

**Endpoint:** `DELETE /users/:id`

**Description:** Deletes a user by its ID.

**Response:**

- **200 OK**

  ```json
  {
    "message": "User deleted"
  }
  ```

- **404 Not Found**

  ```json
  {
    "error": "User not found"
  }
  ```


This `api.md` file provides a comprehensive overview of the available API endpoints, their request bodies, and possible responses for the Eventful Backend Server. It covers authentication, events, roles, teams, and users.
