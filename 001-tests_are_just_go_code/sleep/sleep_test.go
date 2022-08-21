package sleep

import (
	"testing"
	"time"
)

// After starting `go test -v ./...`, run `ps a -u "${USER}" | grep [s]leep` to find this:
// 17928 pts/13   Sl+    0:00 /tmp/go-build886464367/b087/sleep.test -test.testlogfile=/tmp/go-build886464367/b087/testlog.txt -test.paniconexit0 -test.timeout=10m0s -test.v=true
func TestTmpExecutable(t *testing.T) {
	time.Sleep(time.Minute)
}
