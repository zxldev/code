package code

import "testing"

func TestNumDistinct(t *testing.T) {
	r := numDistinct("babgbag", "bag")
	if r != 5 {
		t.Error("失败")
	}
}
