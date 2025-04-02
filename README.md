# Gauth - Go Authentication and Authorization Library

Gauth is a production-grade, secure authentication and authorization library for Go applications. It provides a comprehensive suite of features for handling user authentication, session management, and access control.

## Features

- JWT-based authentication with HTTP-only cookies
- Refresh token management with Redis backend
- Role-based and permission-based access control
- Multi-factor authentication (TOTP)
- OAuth and social login integrations
- Email verification
- Secure password management
- Audit logging for security events
- Token revocation and blacklisting
- Configurable for various SQL databases

## Installation

```bash
go get github.com/amosehiguese/gauth
```

## Quick Start

```go
package main

import (
    "github.com/amosehiguese/gauth"
    "github.com/amosehiguese/gauth/config"
    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize gauth with configurations
    auth := gauth.New(config.Config{
        JWTSecret:          "your-jwt-secret",
        AccessTokenExpiry:  3600,      // 1 hour in seconds
        RefreshTokenExpiry: 2592000,   // 30 days in seconds
        RedisURL:           "redis://localhost:6379",
        DatabaseURL:        "postgres://user:password@localhost:5432/authdb",
    })

    // Initialize Gin router
    router := gin.Default()

    // Set up auth routes
    authGroup := router.Group("/auth")
    {
        authGroup.POST("/register", auth.RegisterHandler)
        authGroup.POST("/login", auth.LoginHandler)
        authGroup.POST("/logout", auth.LogoutHandler)
    }

    // Protected route example
    protectedGroup := router.Group("/api")
    protectedGroup.Use(auth.RequireAuth())
    {
        protectedGroup.GET("/protected", yourProtectedHandler)
    }

    // Start the server
    router.Run(":8080")
}

func yourProtectedHandler(c *gin.Context) {
    // Get user from context
    user, _ := c.Get("user")
    c.JSON(200, gin.H{"message": "Authenticated!", "user": user})
}
```

## Documentation

For detailed documentation, code examples, and best practices, visit [docs/README.md](docs/README.md).

## Security

Gauth follows OWASP security standards and best practices for authentication systems. 

## License

Copyright 2025 gauth

Licensed under the MIT License. <br/> See [LICENSE.md](LICENSE.md) for more information.



## Contributors ✨

<a href="https://github.com/amosehiguese/gauth/graphs/contributors">
  <p align="left">
    <img width="220" src="https://contrib.rocks/image?repo=amosehiguese/gauth" alt="A table of avatars from the project's contributors" />
  </p>
</a>

## Star History

Truly grateful for your support 💖