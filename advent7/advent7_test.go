package advent7

import (
	"testing"
)

func TestSplitTargetOperands(t *testing.T) {

	target, operands, err := SplitTargetOperands("127:31 20 13 15")
	if err != nil {
		t.Fatalf("Unexpected error parsing numbers")
	}
	t.Errorf("target: %d, operands: %v", target, operands)
}
