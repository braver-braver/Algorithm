package dp_test

import "testing"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var res int
	var f1, f2 = 0, 1
	for i := 2; i <= n; i++ {
		res = f2 + f1
		f1, f2 = f2, res
	}
	return res
}

func TestFib(t *testing.T) {
	t.Log(fib(2))
	t.Log(fib(4))
}
