package requests

import (
	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/http/validation"
)

// Create task form request
func CreateTask() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"title":       validation.Rules{"required", "min:5", "max:120"},
		"description": validation.Rules{"required", "min:5"},
	})
}

// Update task form request
func UpdateTask() http.Middleware {
	return http.FormRequest(validation.RulesBag{
		"title":       validation.Rules{"required", "min:5", "max:120"},
		"description": validation.Rules{"required", "min:5"},
	})
}
