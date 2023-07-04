package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/api/public/register", register)
	http.HandleFunc("/api/public/login", login)
	http.HandleFunc("/api/private/self", self)

	http.HandleFunc("/api/public/log/register", LogWrapper(register))
	http.HandleFunc("/api/public/log/login", LogWrapper(login))
	http.HandleFunc("/api/private/log/self", LogWrapper(self))

	fmt.Println("Server is listening on port 8090")
	http.ListenAndServe(":8090", nil)

}

/*
	TODO #2: ✅
	- implement the logic to register a new user (username, password, full_name, address)
	- Validate username (not empty and unique)
	- Validate password (length should at least 8)
*/
func register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	fmt.Println("Register handler")

	// Parse the JSON request body into req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	// fmt.Println("Register request:", req)

	if len(req.Username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username is required")
		return
	}

	if _, err := userStore.Get(req.Username); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username existed")
		return
	}

	if len(req.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Password must be at least 8 characters")
		return
	}

	// Create a new UserInfo object
	user := UserInfo{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Address:  req.Address,
	}

	if err := userStore.Save(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error:", err.Error())
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User %s created", req.Username)
	return
}

type RegisterRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	Address   string `json:"address"`
}

/*
	TODO #3: ✅
	- implement the logic to login
	- validate the user's credentials (username, password)
	- Return JWT token to client
*/
func login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	jsonErr := json.NewDecoder(r.Body).Decode(&req)
	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", jsonErr.Error())
		return
	}

	user, err := userStore.Get(req.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid username or password")
		return
	}

	if user.Password != req.Password {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid username or password")
		return
	}

	token, err := GenerateToken(user.Username, 24*60*60*time.Second)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	resp := LoginResponse{Token: token}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp.Token))
	return
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
}

/*
TODO #4: ✅
- implement the logic to get user info
- Extract the JWT token from the header
- Validate Token
- Return user info`
*/
func self(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	fmt.Println("Self handler", authHeader)
	extractUserNameFn := func(authenticationHeader string) (string, error) {
		tokenString := extractTokenFromHeader(authenticationHeader)
		fmt.Println("Token:", tokenString)
		subject, err := ValidateJWTToken(tokenString)
		fmt.Println("Subject:", subject)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return "", err
		}
		return subject, nil
	}

	username, err := extractUserNameFn(authHeader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	user, _ := userStore.Get(username)
	jsonBytes, jsonErr := json.Marshal(user)

	if jsonErr != nil || len(user.Username) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: getting user info")
		return
	}

	fmt.Println("User:", user)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return
}

func extractTokenFromHeader(authHeader string) string {
	// Assuming the token is present in the "Bearer" scheme
	const bearerPrefix = "Bearer "

	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):]
	}

	return ""
}

/*
TODO: extra wrapper
Print some logs to console
  - Path
  - Http Status code
  - Time start, Duration
*/
func LogWrapper(handler http.HandlerFunc) http.HandlerFunc {
	// panic("TODO implement me")
	return handler
}

/*
	TODO #1: implement in-memory user store
	TODO #2: implement register handler
	TODO #3: implement login handler
	TODO #4: implement self handler

	Extra: implement log handler
*/
