package util_test

import (
	"testing"

	"github.com/vathsalashetty96/argo-cd/util"
)

func TestMakeSignature(t *testing.T) {
	for size := 1; size <= 64; size++ {
		s, err := util.MakeSignature(size)
		if err != nil {
			t.Errorf("Could not generate signature of size %d: %v", size, err)
		}
		t.Logf("Generated token: %v", s)
	}
}
