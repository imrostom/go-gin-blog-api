package helpers

import (
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/imrostom/go-blog-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	// Hash the password with bcrypt using the default cost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Return the hashed password as a string
	return string(hashedPassword), nil
}

// VerifyHashPassword compares the hashed password with the plain text password
func VerifyHashPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Define your signing key (keep it secret)
var jwtSecret = []byte("mySecretJwtToken")

// Define a struct to hold JWT claims
type Claims struct {
	Data map[string]string `json:"data"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(user *models.User) (string, error) {
	expirationTime := time.Now().Add(60 * 24 * 7 * time.Minute)

	// Create the claims, which includes the email and expiry time
	claims := &Claims{
		Data: map[string]string{
			"id":    strconv.FormatUint(uint64(user.ID), 10),
			"name":  user.Name,
			"email": user.Email,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create a new token using the signing method (HS256) and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT parses and validates a JWT token
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse the token with the claims and the secret key
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GenerateSlug(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

// UploadImage handles the logic for uploading an image file with a unique name
func UploadImage(c *gin.Context, file *multipart.FileHeader, uploadDir string) (string, error) {
	// Ensure the upload directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, os.ModePerm) // Create directory if it doesn't exist
	}

	// Generate a unique file name with the original extension
	uniqueFileName := uuid.New().String() + filepath.Ext(file.Filename)

	// Construct the full file path
	filePath := filepath.Join(uploadDir, uniqueFileName)

	// Save the uploaded file to the specified location
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", errors.New("failed to save image: " + err.Error())
	}

	return filePath, nil
}
