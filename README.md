# Role Based Access Boilerplate In Go

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
