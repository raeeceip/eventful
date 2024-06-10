# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.1] - 2024-06-10
### Added
- Initial release of the Eventful Backend Server.
- JWT authentication using golang-jwt/jwt/v5.
- CRUD operations for events, roles, teams, and users.
Add handler and repository tests

- Added unit tests for User, Role, Team, and Event handlers.
- Implemented test cases for Create, Get by ID, Update, and Delete operations.
- Configured test setup with in-memory SQLite database.
- Verified functionality of authentication middleware and token generation.
- Ensured type consistency and proper error handling in repository functions.
