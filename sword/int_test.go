package sword

import (
	"reflect"
	"testing"
)

func Test_binaryExample(t *testing.T) {
	binaryExample()
}

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				dividend: 10,
				divisor:  3,
			},
			want: 3,
		},
		{
			args: args{
				dividend: 7,
				divisor:  -3,
			},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{n: 2},
			want: []int{0, 1, 1},
		},
		{
			args: args{n: 5},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBits2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{n: 2},
			want: []int{0, 1, 1},
		},
		{
			args: args{n: 5},
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			want: 16,
		},
		{
			args: args{words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			want: 4,
		},
		{
			args: args{words: []string{"a", "aa", "aaa", "aaaa"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.words); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProduct2(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}},
			want: 16,
		},
		{
			args: args{words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}},
			want: 4,
		},
		{
			args: args{words: []string{"a", "aa", "aaa", "aaaa"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct2(tt.args.words); got != tt.want {
				t.Errorf("maxProduct2() = %v, want %v", got, tt.want)
			}
		})
	}
}
