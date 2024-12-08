package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"project-crud_baru/controllers/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func VerifyJWT(token, secret string) (map[string]interface{}, bool) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, false
	}

	header, payload, signature := parts[0], parts[1], parts[2]

	// Verify signature
	expectedSignature := utils.CreateSignature(header, payload, secret)
	if signature != expectedSignature {
		return nil, false
	}

	// Decode payload
	payloadJSON, _ := base64.RawURLEncoding.DecodeString(payload)
	var payloadData map[string]interface{}
	if err := json.Unmarshal(payloadJSON, &payloadData); err != nil {
		return nil, false
	}

	// Check expiration
	if exp, ok := payloadData["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return nil, false
		}
	}

	return payloadData, true
}

// AuthMiddleware checks the validity of the JWT token
func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
	}

	secret := "your_secret_key"
	_, valid := VerifyJWT(token, secret)
	if !valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	return c.Next()
}

// Middleware to check role
func CheckRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			// Decode token (assuming it's base64-encoded JSON)
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}
			token := parts[1]
			data, err := base64.StdEncoding.DecodeString(token)
			if err != nil {
				http.Error(w, "Failed to decode token", http.StatusUnauthorized)
				return
			}

			// Parse token data
			var payload map[string]interface{}
			err = json.Unmarshal(data, &payload)
			if err != nil {
				http.Error(w, "Failed to parse token", http.StatusUnauthorized)
				return
			}

			// Check role
			if payload["role"] != role {
				http.Error(w, "Unauthorized role", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// Middleware to check jenis_user
func CheckJenisUser(ju string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header missing", http.StatusUnauthorized)
				return
			}

			// Decode token (assuming it's base64-encoded JSON)
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}
			token := parts[1]
			data, err := base64.StdEncoding.DecodeString(token)
			if err != nil {
				http.Error(w, "Failed to decode token", http.StatusUnauthorized)
				return
			}

			// Parse token data
			var payload map[string]interface{}
			err = json.Unmarshal(data, &payload)
			if err != nil {
				http.Error(w, "Failed to parse token", http.StatusUnauthorized)
				return
			}

			// Check jenis_user
			if payload["jenis_user"] != ju {
				http.Error(w, "Unauthorized jenis_user", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}