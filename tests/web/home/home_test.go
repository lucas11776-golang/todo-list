package home

import (
	"server/tests"
	"testing"

	"github.com/lucas11776-golang/http"
)

func xTestHomePage(t *testing.T) {
	res := tests.TestCase(t).
		Request().
		Get("/")

	res.AssertStatusCode(http.HTTP_RESPONSE_OK).
		AssertIsView().
		AssertView("index")
}

func TestLoginPage(t *testing.T) {
	// TODO problem CWD is test execution path :(...
	res := tests.TestCase(t).
		Request().
		Get("login")

	res.AssertOk().
		AssertIsView().
		AssertView("index")
}
