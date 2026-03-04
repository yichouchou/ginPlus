package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// HeaderCheckResult represents the result of header validation
type HeaderCheckResult struct {
	Valid   bool
	Message string
	Headers map[string]string
}

// RequiredHeaders defines common required headers
var RequiredHeaders = []string{
	"Content-Type",
	"Accept",
}

// HeaderValidator validates request headers
type HeaderValidator struct {
	Required []string
	Optional []string
}

// NewHeaderValidator creates a new header validator
func NewHeaderValidator(required, optional []string) *HeaderValidator {
	return &HeaderValidator{
		Required: required,
		Optional: optional,
	}
}

// ValidateHeaders validates the required headers are present
func (hv *HeaderValidator) ValidateHeaders(c *gin.Context) *HeaderCheckResult {
	result := &HeaderCheckResult{
		Valid:   true,
		Headers: make(map[string]string),
	}

	// Check required headers
	for _, header := range hv.Required {
		value := c.GetHeader(header)
		if value == "" {
			result.Valid = false
			result.Message = "Missing required header: " + header
			return result
		}
		result.Headers[header] = value
	}

	// Collect optional headers
	for _, header := range hv.Optional {
		if value := c.GetHeader(header); value != "" {
			result.Headers[header] = value
		}
	}

	return result
}

// CheckContentType validates the Content-Type header
func CheckContentType(c *gin.Context, allowedTypes []string) bool {
	contentType := c.ContentType()
	for _, t := range allowedTypes {
		if strings.Contains(contentType, t) {
			return true
		}
	}
	return false
}

// GetHeaderMap converts gin headers to a map
func GetHeaderMap(c *gin.Context) map[string]string {
	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	return headers
}

// HasHeader checks if a specific header exists
func HasHeader(c *gin.Context, header string) bool {
	return c.GetHeader(header) != ""
}

// GetHeaderValue gets a header value with fallback
func GetHeaderValue(c *gin.Context, header, fallback string) string {
	if value := c.GetHeader(header); value != "" {
		return value
	}
	return fallback
}
