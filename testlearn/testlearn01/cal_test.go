package testlearn01

import "testing"

func TestCal(t *testing.T) {
	res := cal(10)
	if res != 55 {
		t.Fatalf("Cal(10)执行错误，期望值:%v，实际值:%v", 55, res)
	} else {
		t.Fatalf("Cal(10)执行正确，期望值:%v，实际值:%v", 55, res)
	}
}
