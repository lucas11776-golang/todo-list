package tasks

import (
	"server/models"

	"github.com/lucas11776-golang/http"
	"github.com/lucas11776-golang/orm"
	"github.com/spf13/cast"
)

// Check if user is guest - not logged int
func UserTask(req *http.Request, res *http.Response, next http.Next) *http.Response {
	count, err := orm.Model(models.Task{}).
		Where("id", "=", cast.ToInt64(req.Parameters.Get("task"))).
		AndWhere("user_id", "=", cast.ToInt64(req.Session.Get("user_id"))).
		Count()

	if err != nil || count == 0 {
		return res.Redirect("todos")
	}

	return next()
}
