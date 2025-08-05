package requests

import (
	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/validation"
)

// Login Form Validation
func LoginFormRequest() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"email":    validation.Rules{"required", "string", "email", "exists:users,default"},
		"password": validation.Rules{"required", "string"},
	})
}

// Register Form Validation
func RegisterFormRequest() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"first_name": validation.Rules{"required", "string", "min:3", "max:30"},
		"last_name":  validation.Rules{"required", "string", "min:3", "max:30"},
		"email":      validation.Rules{"required", "string", "email", "unique:users,default"},
		"password":   validation.Rules{"required", "string", "min:8", "max:21", "confirmed"},
	})
}
