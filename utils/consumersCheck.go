package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// ConsumersCheck validates the Accept header against supported content types
type ConsumersCheck struct {
	SupportedTypes []string
}

// NewConsumersCheck creates a new consumers checker
func NewConsumersCheck(supportedTypes []string) *ConsumersCheck {
	return &ConsumersCheck{
		SupportedTypes: supportedTypes,
	}
}

// Check verifies if the request accepts a supported content type
func (cc *ConsumersCheck) Check(c *gin.Context) bool {
	accept := c.GetHeader("Accept")
	if accept == "" {
		return true // No Accept header means accept all
	}

	for _, supported := range cc.SupportedTypes {
		if strings.Contains(accept, supported) {
			return true
		}
	}
	return false
}

// GetAcceptedType returns the best matching content type
func (cc *ConsumersCheck) GetAcceptedType(c *gin.Context) string {
	accept := c.GetHeader("Accept")
	if accept == "" {
		if len(cc.SupportedTypes) > 0 {
			return cc.SupportedTypes[0]
		}
		return "*/*"
	}

	for _, supported := range cc.SupportedTypes {
		if strings.Contains(accept, supported) {
			return supported
		}
	}
	return ""
}

// CheckProduce validates the Content-Type header
func CheckProduce(c *gin.Context, supportedTypes []string) bool {
	contentType := c.ContentType()
	for _, supported := range supportedTypes {
		if strings.Contains(contentType, supported) {
			return true
		}
	}
	return false
}
