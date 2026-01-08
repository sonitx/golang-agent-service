package utils

import (
	"errors"
	"strings"
)

// write function to validate SQL query, only allow select query
func ValidateSQLQuery(query string) error {
	if strings.Contains(query, "INSERT") || strings.Contains(query, "UPDATE") || strings.Contains(query, "DELETE") {
		return errors.New("invalid query")
	}
	return nil
}
