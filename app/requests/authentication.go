package requests

import (
	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/validation"
)

// Login Form Validation
func LoginFormRequest() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"email":    validation.Rules{"required"},
		"password": validation.Rules{"required"},
	})
}

// Login Form Validation
func RegisterFormRequest() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"first_name": validation.Rules{"required"},
		"last_name":  validation.Rules{"required"},
		"email":      validation.Rules{"required"},
		"password":   validation.Rules{"required", "min:8", "max:21", "confirmed"},
	})
}
