package greedy

import (
	"strconv"
	"testing"
)

func monotoneIncreasingDigits(n int) int {
	var str = strconv.Itoa(n)
	ss := []byte(str)
	var l = len(ss)
	for i := l - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1]--
			for j := i; j < l; j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}

func Test_monotoneIncreasingDigits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				10,
			},
			9,
		},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.n); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
