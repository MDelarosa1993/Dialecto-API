package handlers

import (
	"bytes"
	"Dialecto-API/db"
	"Dialecto-API/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
	return router
}


func TestRegisterUser(t *testing.T) {
	os.Setenv("DATABASE_URL", "host=localhost user=postgres password=yourpassword dbname=dialecto port=5432 sslmode=disable")
	db.ConnectDB()
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	router := setupRouter()

	regData := map[string]string{
		"email":    "testregister@example.com",
		"password": "password123",
	}
	regBody, err := json.Marshal(regData)
	if err != nil {
		t.Fatalf("Failed to marshal registration data: %v", err)
	}

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(regBody))
	if err != nil {
		t.Fatalf("Failed to create registration request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
	}
}

func TestLoginUser(t *testing.T) {
	os.Setenv("DATABASE_URL", "host=localhost user=postgres password=yourpassword dbname=dialecto port=5432 sslmode=disable")
	db.ConnectDB()
	
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	router := setupRouter()

	regData := map[string]string{
		"email":    "testlogin@example.com",
		"password": "password123",
	}
	regBody, err := json.Marshal(regData)
	if err != nil {
		t.Fatalf("Failed to marshal registration data: %v", err)
	}
	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(regBody))
	if err != nil {
		t.Fatalf("Failed to create registration request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	loginBody, err := json.Marshal(regData)
	if err != nil {
		t.Fatalf("Failed to marshal login data: %v", err)
	}
	req, err = http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
	if err != nil {
		t.Fatalf("Failed to create login request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal login response: %v", err)
	}

	token, exists := response["token"]
	if !exists {
		t.Error("Expected token in response, but it was missing")
	}
	if token == "" {
		t.Error("Expected non-empty token, but got an empty string")
	}
}
