package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ProducesSet defines the content types that a handler can produce
type ProducesSet struct {
	ContentTypes []string
}

// NewProducesSet creates a new produces set
func NewProducesSet(contentTypes []string) *ProducesSet {
	return &ProducesSet{
		ContentTypes: contentTypes,
	}
}

// SetContentType sets the response Content-Type header
func (ps *ProducesSet) SetContentType(c *gin.Context) {
	if len(ps.ContentTypes) > 0 {
		c.Header("Content-Type", ps.ContentTypes[0])
	}
}

// GetContentType returns the appropriate content type based on Accept header
func (ps *ProducesSet) GetContentType(c *gin.Context) string {
	accept := c.GetHeader("Accept")
	
	// If no Accept header, use first supported type
	if accept == "" {
		if len(ps.ContentTypes) > 0 {
			return ps.ContentTypes[0]
		}
		return "application/json"
	}

	// Try to match Accept header with supported types
	for _, ct := range ps.ContentTypes {
		if strings.Contains(accept, ct) || accept == "*/*" {
			return ct
		}
	}

	// Default to first supported type
	if len(ps.ContentTypes) > 0 {
		return ps.ContentTypes[0]
	}
	return "application/json"
}

// Common content types
const (
	ContentTypeJSON = "application/json"
	ContentTypeXML  = "application/xml"
	ContentTypeHTML = "text/html"
	ContentTypeText = "text/plain"
	ContentTypeForm = "application/x-www-form-urlencoded"
	ContentTypeMulti = "multipart/form-data"
)
