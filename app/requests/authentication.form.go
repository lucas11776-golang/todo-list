package requests

import (
	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/validation"
)

// Login Form Validation
func LoginFormRequest() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"email":    validation.Rules{"required", "string", "email", "exists:users,sqlite"},
		"password": validation.Rules{"required", "string"},
	})
}

// Login Form Validation
func RegisterFormRequest() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"first_name": validation.Rules{"required", "string", "min:3", "max:30"},
		"last_name":  validation.Rules{"required", "string", "min:3", "max:30"},
		"email":      validation.Rules{"required", "string", "email", "unique:users,sqlite"},
		"password":   validation.Rules{"required", "string", "min:8", "max:21", "confirmed"},
	})
}
