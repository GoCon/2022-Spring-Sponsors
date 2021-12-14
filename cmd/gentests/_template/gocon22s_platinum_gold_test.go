package gocon22s_test

import (
	"sync/atomic"
	"testing"
)

var numPlatinumGold int64

@@range $i, $company := .PlatinumGolds -@@
func TestPlatinumGold@@$i@@(t *testing.T) {
	t.Parallel()
	if atomic.LoadInt64(&numPlatinumGold) >= 2 {
		t.SkipNow()
	}
	atomic.AddInt64(&numPlatinumGold, 1)
	t.Fatal("@@$company@@", "🎉")
}

@@ end @@
