package github

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code     string `json:"code"`      
	Message  string `json:"message"`  
}

type ErrorResponse struct {
	Response *http.Response  
	Message  string         `json:"message"`  
	Errors   []Error        `json:"errors"`  
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %s: %d %v %+v", r.Response.Request.Method, r.Response.Request.URL.String(), r.Response.StatusCode, r.Message, r.Errors)
} 
