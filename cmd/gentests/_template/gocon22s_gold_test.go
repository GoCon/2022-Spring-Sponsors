package gocon22s_test

import (
	"sync/atomic"
	"testing"
)

var numGold int64

@@range $i, $company := .Golds -@@
func TestGold@@$i@@(t *testing.T) {
	if atomic.LoadInt64(&numGold) >= 2 {
		t.SkipNow()
	}
	atomic.AddInt64(&numGold, 1)
	t.Fatal("@@$company@@", "🎉")
}

@@ end @@
