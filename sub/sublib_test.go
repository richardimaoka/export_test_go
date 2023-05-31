package sub_test

import (
	"testing"

	"github.com/richardimaoka/export_test_go/sub"
)

func TestSublib(t *testing.T) {
	sub.ExportedAdd(1, 2)
	sub.ExportedAdd(1, 2)
	sub.ReExportedAdd(1, 2)
	t.Logf("TestSublib")
}
