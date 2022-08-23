package random

import (
	"math/rand"
	"testing"
	"time"
)

func TestPick(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("seed = %d", seed)
	r := rand.New(rand.NewSource(seed))

	data := make([]int, 5)
	for i := 0; i < 5; i++ {
		data[i] = r.Int()
	}
	t.Logf("data[] = %#v", data)

	got := pick(data)
	for _, v := range data {
		if got == v {
			return
		}
	}
	t.Errorf("pick(%v) = %d; not in slice", data, got)
}
