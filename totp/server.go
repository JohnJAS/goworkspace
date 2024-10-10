package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

type KeyRequest struct {
	UserName string `json:"username"`
}

type KeyResponse struct {
	URL    string `json:"url"`
	Secret string `json:"secret"`  // 生产环境中请不要直接返回密钥
	QRCode string `json:"qr_code"` // 返回二维码的 Base64 字符串
}

type ValidateRequest struct {
	UserName string `json:"username"`
	Code     string `json:"code"`
}

type ValidateResponse struct {
	Valid bool `json:"valid"`
}

// 用户密钥存储
var userSecrets = make(map[string]string)

func generateKey(w http.ResponseWriter, r *http.Request) {
	var req KeyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: req.UserName,
	})
	if err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}
	fmt.Println("Key: ", key)

	// 存储用户的密钥
	userSecrets[req.UserName] = key.Secret()

	// 生成二维码
	qrCode, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Error generating QR code", http.StatusInternalServerError)
		return
	}

	// 将二维码转换为 Base64 字符串
	qrCodeBase64 := base64.StdEncoding.EncodeToString(qrCode)

	response := KeyResponse{
		URL:    key.URL(),
		Secret: key.Secret(), // 生产环境中请不要直接返回密钥
		QRCode: qrCodeBase64,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func generateQRImage(w http.ResponseWriter, r *http.Request) {
	var req KeyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MyApp",
		AccountName: req.UserName,
	})
	if err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}

	// 存储用户的密钥
	userSecrets[req.UserName] = key.Secret()

	// 生成二维码图像
	qrCode, err := qrcode.Encode(key.URL(), qrcode.Medium, 256)
	if err != nil {
		http.Error(w, "Error generating QR code", http.StatusInternalServerError)
		return
	}

	// 设置响应头为 image/png
	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	// 将二维码图像写入响应
	w.Write(qrCode)
}

func validateCode(w http.ResponseWriter, r *http.Request) {
	var req ValidateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 根据用户名获取密钥
	secret, exists := userSecrets[req.UserName]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	valid := totp.Validate(req.Code, secret)
	response := ValidateResponse{Valid: valid}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/generate-key", generateKey)
	http.HandleFunc("/generate-image", generateQRImage)
	http.HandleFunc("/validate-code", validateCode)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
