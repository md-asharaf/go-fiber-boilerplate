package utils

import (
	"regexp"
	"strings"
)

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
