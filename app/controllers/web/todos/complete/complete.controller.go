package complete

import (
	"server/app/services/tasks"

	"github.com/lucas11776-golang/http"
	"github.com/spf13/cast"
)

// Store Todo Item
func Store(req *http.Request, res *http.Response) *http.Response {
	err := tasks.TaskCompleted(cast.ToInt64(req.Parameters.Get("task")))

	if err != nil {
		// TODO: log error
	}

	return res.Back()
}

// Store Todo Item
func Destroy(req *http.Request, res *http.Response) *http.Response {
	err := tasks.TaskUncompleted(cast.ToInt64(req.Parameters.Get("task")))

	if err != nil {
		// TODO: log error
	}

	return res.Back()
}
