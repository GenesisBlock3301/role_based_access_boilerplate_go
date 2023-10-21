# Role-Based Authentication in GoLang Application

## Overview
This GoLang application implements a role-based authentication system, providing secure access control to various resources. It encompasses features like Google OAuth, JWT token authentication, and email verifications for user management.
## Table Structure
The application's database schema includes the following tables:

**User:** Stores user information including username, email, password hash, and associated roles.
**Role:** Contains roles with corresponding permissions.<br>
**Permission:** Lists specific actions that can be performed within the application.

## Authentication Methods
**1. Google OAuth**
   This application integrates with Google OAuth for seamless user authentication. Users can log in using their Google credentials, which are then used to create or update their profile in the application's database.

**2. JWT Token Authentication**
   JWT tokens are employed for secure communication between the client and the server. After successful authentication, a JWT token is generated and provided to the client, which is then included in subsequent requests for authorization.

**3. Email Verification**
   To ensure a valid email address, users receive a verification link upon registration. Clicking this link confirms the email address and activates the account.

## Role-Based Access Control (RBAC)
Roles play a central role in defining what actions a user can perform within the application. Each role is associated with specific permissions, providing granular control over the features and resources accessible to users.

### Project Structure:
```
role_based_access_boilerplate_go/
├── cmd/
│   └── main.go
├── internal/
│   ├── configurations/
│   │     ├── db/
│   │     │   └── config.go
│   │     │   └── models.go
│   ├── controller/
│   │     ├── user/
│   │     │     └── user-controller.go
│   │     ├── role/
│   │     │     └── roles-controller.go
│   │     └── router.go
│   ├── middleware/
│   │    └── auth-middleware.go
│   ├── routes/
│   │     ├── user/
│   │     │     └── user-router.go
│   │     ├── role/
│   │     │     └── roles-router.go
│   │     └── router.go
│   ├── serializers/
│   │    └── user-serializers.go
│   ├── service/
│   │     ├── user/
│   │     │     └── user-service.go
│   │     ├── role/
│   │     │     └── roles-service.go
├── test/
│   └── my_service_test.go
└── README.md
```

## Explanation:
**/cmd:** Contains the main entry point for your application.
**/internal:** Contains internal packages, including database operations and data models. 
  - **/configurations**: Houses API handlers for user, role, and permission management.
  - **/controllers**: 
  - **/middlewares**:
  - **/routes**:
  - **/schemas**:
  - **/serializers**:
  - **/services**
  - **/utils**:
README.md: This file, providing an overview of the project.

## Getting Started
1. Clone the repository.
2. Initialize the database using the provided SQL script in /scripts.
3. Customize the authentication methods and API handlers according to your requirements.
4. Run the application using go run cmd/your-app/main.go.

## Creating Project: 
- Open project folder
- Command: `go mod init`

## Design database model & making operations:
- Follow GORM orm documentation. 

## How JWT Token Generate:

---
1. Authenticated User login using credentials (username, password).
2. Verify credential of users(Username & password).
3. Add JWT claims. For example: `Authorization`, `exp`, `user_id` etc.
4. Using these claims generate token with secret key.


## How to make JWT token validation when I make call with this Token(middleware):

---
1. Firstly token parse which is provided by frontend and find exact token.
2. Then check token validation, if valid then move forward.
3. From this token, after decoding we can get userID.


## How does Register email verification work:

---
1. After save user's information.
2. Then make email template with verification link and send it on behalf of system to client.
3. Click on that link & activate this user.
