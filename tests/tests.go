package tests

import (
	"os"
	"server/bootstrap"
	"server/database/migration"
	test "testing"

	"github.com/lucas11776-golang/http/testing"
	"github.com/lucas11776-golang/http/utils/strings"
	"github.com/lucas11776-golang/orm"
)

// Comment
func TestCase(t *test.T) *testing.TestCase {
	server := bootstrap.Boot("")

	os.Setenv("APP_URL", "")
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "")
	os.Setenv("DATABASE", ":memory:")
	os.Setenv("SESSION_NAME", "session")
	os.Setenv("SESSION_KEY", strings.Random(10))

	if err := orm.DB.Database("sqlite").Migration().Migrate(migration.Tables()); err != nil {
		t.Fatal(err)
	}

	return testing.NewTestCase(t, server, false)
}
