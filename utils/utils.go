package utils

import (
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GenComment holds the parsed annotation information for a route
type GenComment struct {
	Parms      []*Parm
	Result     []*Parm
	Methods    []string
	RouterPath string
	Note       string
}

// Parm represents a parameter in the route annotation
type Parm struct {
	ParmName       string
	ParmType      reflect.Type
	ParmKind      reflect.Kind
	IsHeaderOrBody HeaderOrBody
	IsMust        bool
	Value         reflect.Value
	ParmKindStr   string
	NewValueStr   string
	StrInTypeOf   string
}

// HeaderOrBody indicates where the parameter comes from
type HeaderOrBody int

const (
	// Default means not specified
	Default HeaderOrBody = iota
	// Header means parameter comes from request header
	Header
	// Body means parameter comes from request body
	Body
)

// GenInfo holds the generated router information
type GenInfo struct {
	Tm             int64
	List           []GenRouterInfo
	PkgImportList   map[string]string
	PkgImportStrs  []string
	PkgName        string
}

// GenRouterInfo holds the router information for a handler
type GenRouterInfo {
	HandFunName string
	GenComment  GenComment
}

// GenInfoCnf is the global configuration for code generation
var GenInfoCnf GenInfo

// ParameterAnnotationRules defines the annotation format for parameters
// Format: [!name string, !password string, !age int]
// The ! prefix indicates a required parameter
// @GET /block
const (
	// ParameterAnnotationRule describes the expected format for parameter annotations
	// [paramName paramType, ...]
	ParameterAnnotationRule = "Parameter annotation format: [name type, ...]"
	
	// RequiredParamPrefix marks a parameter as required
	RequiredParamPrefix = "!"
)

// GetRouter is the default router file name
const GetRouter = "routers/gen_router.go"

// GenTemp is the template for code generation
const GenTemp = `{{range .List}}
{{$genComment := .GenComment}}
// {{.HandFunName}} 
func {{.HandFunName}}(c *gin.Context) {
	{{range $i, $parm := $genComment.Parms}}
	{{$parm.NewValueStr}}
	{{end}}
	{{range $i, $result := $genComment.Result}}
	{{$result.NewResultStr}}
	{{end}}
}
{{end}}`

// ContainsHttpMethod checks if the annotation contains HTTP method
// Returns the HTTP method and whether it was found
func ContainsHttpMethod(annoDoc string) (httpMethod string, isFound bool) {
	upperDoc := strings.ToUpper(annoDoc)
	
	if strings.Contains(upperDoc, "GET") && !strings.Contains(upperDoc, "POST") {
		return "GET", true
	}
	if strings.Contains(upperDoc, "POST") {
		return "POST", true
	}
	if strings.Contains(upperDoc, "DELETE") {
		return "DELETE", true
	}
	if strings.Contains(upperDoc, "PATCH") {
		return "PATCH", true
	}
	if strings.Contains(upperDoc, "PUT") {
		return "PUT", true
	}
	if strings.Contains(upperDoc, "OPTIONS") {
		return "OPTIONS", true
	}
	if strings.Contains(upperDoc, "HEAD") {
		return "HEAD", true
	}
	return "ANY", false
}

// ContainsHttpRouter checks if the annotation contains a route path
func ContainsHttpRouter(annoDoc string) (router string, contains bool) {
	if strings.Contains(annoDoc, "/") {
		indexStart := strings.Index(annoDoc, "/")
		return annoDoc[indexStart:], true
	}
	return "", false
}

// ContainsParmsOrResults parses parameter and result annotations
// Supports two formats:
// - [] for header parameters: [name string, age int]
// - {} for body parameters: {name string, age int}
func ContainsParmsOrResults(annoDoc string, gc *GenComment) *GenComment {
	// Handle header parameters: [...]
	if strings.HasPrefix(annoDoc, "[") && strings.HasSuffix(annoDoc, "]") {
		trimPrefix := strings.TrimPrefix(annoDoc, "[")
		anno := strings.TrimSuffix(trimPrefix, "]")
		split := strings.Split(anno, ",")
		for index, annoParm := range split {
			if gc.Parms == nil {
				gc.Parms = make([]*Parm, len(split))
			}
			gc.Parms[index] = &Parm{
				ParmName:       parseParamName(annoParm),
				IsHeaderOrBody: Header,
			}
		}
	}

	// Handle body parameters: {...}
	if strings.Contains(annoDoc, "{") && strings.Contains(annoDoc, "}") {
		trimPrefix := strings.TrimPrefix(annoDoc, "{")
		anno := strings.TrimSuffix(trimPrefix, "}")
		split := strings.Split(anno, ",")
		Reverse(&split)
		for index, annoParm := range split {
			if gc.Parms == nil {
				continue
			}
			targetIndex := len(gc.Parms) - index - 1
			if targetIndex >= 0 && targetIndex < len(gc.Parms) {
				gc.Parms[targetIndex].IsHeaderOrBody = Body
			}
		}
	}

	return gc
}

// ContainsBraces checks if the annotation contains parameter braces
func ContainsBraces(annoDoc string) (params string, contains bool) {
	if strings.Contains(annoDoc, "[") && strings.Contains(annoDoc, "]") {
		indexStart := strings.Index(annoDoc, "[")
		indexEnd := strings.Index(annoDoc, "]")
		if indexEnd > indexStart {
			return annoDoc[indexStart : indexEnd+1], true
		}
	}
	return "", false
}

// SplitParms parses parameter annotation into Parm structs
// Format: [!name string, !password string, !age int]
// The ! prefix indicates required parameters
func SplitParms(parmsDoc string) []Parm {
	var parmList []Parm
	split := strings.Split(parmsDoc, ",")
	
	for _, s := range split {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		
		var isRequired bool
		paramName := s
		
		if strings.HasPrefix(s, "!") {
			isRequired = true
			paramName = strings.TrimPrefix(s, "!")
			// Try to extract type after space
			parts := strings.Split(paramName, " ")
			if len(parts) >= 2 {
				paramName = parts[0]
			}
		}
		
		parmList = append(parmList, Parm{
			ParmName: paramName,
			IsMust:   isRequired,
		})
	}
	
	return parmList
}

// Kind2String converts reflect.Kind to string representation
func Kind2String(kind reflect.Kind) string {
	switch kind {
	case reflect.String:
		return "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return "int"
	case reflect.Ptr:
		return "pointer"
	case reflect.Struct:
		return "struct"
	case reflect.Array, reflect.Slice:
		return "array"
	case reflect.Interface:
		return "interface"
	case reflect.Float32, reflect.Float64:
		return "float"
	case reflect.Bool:
		return "bool"
	default:
		return "unknown"
	}
}

// parseParamName extracts the parameter name from annotation
func parseParamName(anno string) string {
	anno = strings.TrimSpace(anno)
	parts := strings.Split(anno, " ")
	if len(parts) > 0 {
		return parts[0]
	}
	return anno
}

// delete_extra_space removes redundant spaces from string
func delete_extra_space(s string) string {
	s = strings.Replace(s, "\t", " ", -1)
	reg := regexp.MustCompile(`\s{2,}`)
	return reg.ReplaceAllString(s, " ")
}

// Reverse reverses a string slice in place
func Reverse(arr *[]string) {
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		(*arr)[i], (*arr)[length-1-i] = (*arr)[length-1-i], (*arr)[i]
	}
}

// MapMerge merges two maps, keeping all keys from both
func MapMerge(first, second map[string]string) map[string]string {
	n := make(map[string]string)
	for i, v := range first {
		n[i] = v
	}
	for j, w := range second {
		if _, ok := n[j]; !ok {
			n[j] = w
		}
	}
	return n
}

// MapMergeMost merges two maps, keeping all keys
func MapMergeMost(first, second map[string]string) map[string]string {
	n := make(map[string]string)
	for i, v := range first {
		n[i] = v
	}
	for j, w := range second {
		n[j] = w
	}
	return n
}

// ReplenishParmsOrResults fills in default values for parameters
// If it's a GET request, all parameters come from header
// If some parameters are annotated, the rest get the opposite type
func ReplenishParmsOrResults(gc *GenComment) {
	if gc == nil || gc.Parms == nil {
		return
	}
	
	// If GET request, all params from header
	if len(gc.Methods) == 1 && gc.Methods[0] == "GET" {
		for _, parm := range gc.Parms {
			parm.IsHeaderOrBody = Header
		}
		return
	}

	// Determine the majority type
	var hasHeader, hasBody bool
	for _, parm := range gc.Parms {
		if parm.IsHeaderOrBody == Header {
			hasHeader = true
		} else if parm.IsHeaderOrBody == Body {
			hasBody = true
		}
	}

	// Fill in defaults
	defaultType := Body
	if hasHeader && !hasBody {
		defaultType = Body
	} else if hasBody && !hasHeader {
		defaultType = Header
	}

	for _, parm := range gc.Parms {
		if parm.IsHeaderOrBody == Default {
			parm.IsHeaderOrBody = defaultType
		}
	}
}

// RandString generates a random string of given length
func RandString(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes) + strconv.Itoa(rand.Int())
}
