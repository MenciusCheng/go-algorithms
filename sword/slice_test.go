package sword

import (
	"reflect"
	"testing"
)

func Test_arrayExample(t *testing.T) {
	arrayExample()
}

func Test_sliceExample(t *testing.T) {
	sliceExample()
}

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2}, {-1, 0, 1},
			},
		},
		{
			args: args{
				nums: []int{},
			},
			want: [][]int{},
		},
		{
			args: args{
				nums: []int{0},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			args: args{
				target: 15,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numSubarrayProductLessThanK(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{10, 5, 2, 6},
				k:    100,
			},
			want: 8,
		},
		{
			args: args{
				nums: []int{1, 2, 3},
				k:    0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubarrayProductLessThanK(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("numSubarrayProductLessThanK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{1, 9, 1},
				k:    2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{0, 1, 0},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{0, 1, 1},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{0, 0, 1, 0, 0, 0, 1, 1},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("findMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 7, 3, 6, 5, 6},
			},
			want: 3,
		},
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: -1,
		},
		{
			args: args{
				nums: []int{2, 1, -1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumMatrix_SumRegion(t *testing.T) {
	/*
		输入:
		["NumMatrix","sumRegion","sumRegion","sumRegion"]
		[[[[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]],[2,1,4,3],[1,1,2,2],[1,2,2,4]]
		输出:
		[null, 8, 11, 12]

		解释:
		NumMatrix numMatrix = new NumMatrix([[3,0,1,4,2],[5,6,3,2,1],[1,2,0,1,5],[4,1,0,1,7],[1,0,3,0,5]]);
		numMatrix.sumRegion(2, 1, 4, 3); // return 8 (红色矩形框的元素总和)
		numMatrix.sumRegion(1, 1, 2, 2); // return 11 (绿色矩形框的元素总和)
		numMatrix.sumRegion(1, 2, 2, 4); // return 12 (蓝色矩形框的元素总和)
	*/
	m := Constructor([][]int{
		{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5},
	})

	type args struct {
		row1 int
		col1 int
		row2 int
		col2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				row1: 2,
				col1: 1,
				row2: 4,
				col2: 3,
			},
			want: 8,
		},
		{
			args: args{
				row1: 1,
				col1: 1,
				row2: 2,
				col2: 2,
			},
			want: 11,
		},
		{
			args: args{
				row1: 1,
				col1: 2,
				row2: 2,
				col2: 4,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := m.SumRegion(tt.args.row1, tt.args.col1, tt.args.row2, tt.args.col2); got != tt.want {
				t.Errorf("SumRegion() = %v, want %v", got, tt.want)
			}
		})
	}
}
