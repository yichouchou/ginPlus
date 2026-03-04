package bind

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

// ParameterBindingSource defines where to bind parameters from
type ParameterBindingSource int

const (
	// BindingFromQuery binds from URL query parameters
	BindingFromQuery ParameterBindingSource = iota
	// BindingFromHeader binds from request headers
	BindingFromHeader
	// BindingFromBody binds from request body (JSON)
	BindingFromBody
	// BindingFromForm binds from form data
	BindingFromForm
	// BindingFromPath binds from URL path parameters
	BindingFromPath
)

// AutoBindParams automatically binds parameters from request to the given struct
// Supports: JSON body, Form data, Query parameters, Headers
func AutoBindParams(c *gin.Context, params interface{}) error {
	return c.ShouldBind(params)
}

// BindParamsFromSource binds params from a specific source
func BindParamsFromSource(c *gin.Context, source ParameterBindingSource, params interface{}) error {
	switch source {
	case BindingFromQuery:
		return c.ShouldBindQuery(params)
	case BindingFromHeader:
		return c.ShouldBindHeader(params)
	case BindingFromBody:
		return c.ShouldBindJSON(params)
	case BindingFromForm:
		return c.ShouldBind(params)
	case BindingFromPath:
		return c.ShouldBindUri(params)
	default:
		return c.ShouldBind(params)
	}
}

// BindByAnnotation binds parameters based on annotation rules
// This function determines the binding method based on HTTP method and parameter types
func BindByAnnotation(c *gin.Context, params []interface{}, method string) ([]interface{}, error) {
	results := make([]interface{}, len(params))

	for i, param := range params {
		paramType := reflect.TypeOf(param)
		if paramType == nil {
			continue
		}

		// Handle pointer types
		if paramType.Kind() == reflect.Ptr {
			paramType = paramType.Elem()
		}

		var err error
		switch method {
		case "GET":
			// For GET requests, try query first, then path
			err = c.ShouldBindQuery(param)
		case "POST", "PUT", "PATCH":
			// For POST/PUT/PATCH, try JSON first, then form
			err = c.ShouldBind(param)
			if err != nil {
				err = c.ShouldBindJSON(param)
			}
		default:
			err = c.ShouldBind(param)
		}

		if err != nil {
			return results, err
		}

		results[i] = param
	}

	return results, nil
}

// BindStruct automatically determines the best binding method based on Content-Type
func BindStruct(c *gin.Context, structPtr interface{}) error {
	contentType := c.ContentType()

	switch contentType {
	case "application/json":
		return c.ShouldBindJSON(structPtr)
	case "application/x-www-form-urlencoded":
		return c.ShouldBind(structPtr)
	default:
		// Try JSON first, then query, then form
		if err := c.ShouldBindJSON(structPtr); err != nil {
			if err := c.ShouldBindQuery(structPtr); err != nil {
				return c.ShouldBind(structPtr)
			}
		}
		return nil
	}
}

// AutoBind is a convenient wrapper that automatically determines
// the best binding method based on the HTTP method and request content type
func AutoBind(c *gin.Context, target interface{}) error {
	method := c.Request.Method

	switch method {
	case "GET":
		// GET requests: bind from query parameters
		return c.ShouldBindQuery(target)
	case "POST", "PUT", "PATCH":
		// POST/PUT/PATCH: determine by Content-Type
		contentType := c.ContentType()
		switch contentType {
		case "application/json":
			return c.ShouldBindJSON(target)
		case "application/x-www-form-urlencoded":
			return c.ShouldBind(target)
		default:
			return c.ShouldBind(target)
		}
	case "DELETE":
		// DELETE: bind from query or body
		if err := c.ShouldBindQuery(target); err != nil {
			return c.ShouldBind(target)
		}
		return nil
	default:
		return c.ShouldBind(target)
	}
}
