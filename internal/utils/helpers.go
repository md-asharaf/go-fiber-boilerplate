package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

// BindAndValidate decodes and validates a request body into the given struct
func BindAndValidate(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	return ValidateStruct(v)
}

// IsValidEmail validates email format
func IsValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// SanitizeString removes extra whitespace and trims string
func SanitizeString(s string) string {
	// Remove extra whitespace
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")

	// Trim leading and trailing whitespace
	return strings.TrimSpace(s)
}

// Contains checks if a slice contains a specific string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
