package math_test

import (
	"testing"

	"github.com/richardimaoka/export_test_go/math"
)

func TestName(t *testing.T) {
	if 2 != math.Sum(1, 1) {
		t.Fatalf("Sum of %d + %d not 2", 1, 1)
	}
}
