package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"crypto/rand"
	"encoding/hex"
)

// Fungsi untuk menghasilkan string acak dengan panjang tertentu
func GenerateRandomString(length int) (string, error) {
    bytes := make([]byte, length/2)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
}

// Fungsi untuk meng-encode data ke base64 tanpa padding
func base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

// Fungsi untuk membuat signature menggunakan HMAC-SHA256
func CreateSignature(header, payload, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(header + "." + payload))
	return base64Encode(h.Sum(nil))
}

// Fungsi untuk generate token JWT
func GenerateJWT(username, role string, idRole string) (string, error) {
	// Header
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	headerJSON, _ := json.Marshal(header)
	headerEncoded := base64Encode(headerJSON)

	// Payload
	payload := map[string]interface{}{
		"username": username,
		"role":     role,
		"id_role":  idRole, // Tambahkan id_role ke payload
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	payloadJSON, _ := json.Marshal(payload)
	payloadEncoded := base64Encode(payloadJSON)

	// Signature
	secret := "your_secret_key"
	signature := CreateSignature(headerEncoded, payloadEncoded, secret)

	// Token JWT
	token := fmt.Sprintf("%s.%s.%s", headerEncoded, payloadEncoded, signature)
	return token, nil
}

// GetError sends an error response

func GetError(err error, w http.ResponseWriter) {

    w.WriteHeader(http.StatusInternalServerError)

    json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

}