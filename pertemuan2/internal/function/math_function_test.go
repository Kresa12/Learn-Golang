package function

import (
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(10, 20)
	expect := 30
	if res != expect {
		t.Errorf("hasil %d harusnya %d", res, expect)
	}
}
