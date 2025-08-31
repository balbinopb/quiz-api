package models

// RegisterRequest represents registration request payload
type RegisterRequest struct {
    Email    string `json:"email" Email:"user@gmail.com"`
    Username string `json:"username" Username:"john_doe"`
    Password string `json:"password" Password:"secret123"`
}

// RegisterResponse represents registration success response
type RegisterResponse struct {
    Message string `json:"message" Message:"user registered successfully"`
}

// LoginRequest represents login request payload
type LoginRequest struct {
    Email    string `json:"email" Email:"user@gmail.com"`
    Password string `json:"password" Password:"secret123"`
}

// LoginResponse represents login response payload
type LoginResponse struct {
    Token string `json:"token" Token:"jwt token "`
}

// ErrorResponse represents standard error
type ErrorResponse struct {
    Error string `json:"error" Error:"invalid email or password"`
}
