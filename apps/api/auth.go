package main

import (
    "encoding/json"
    "net/http"
)

type Credentials struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

var validCredentials = Credentials{
    Name:  "testuser",
    Email: "test@example.com",
}

func authHandler(w http.ResponseWriter, r *http.Request) {
    var creds Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if creds.Name == validCredentials.Name && creds.Email == validCredentials.Email {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Authentication successful"))
    } else {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
    }
}
