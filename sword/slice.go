package sword

import (
	"fmt"
	"math"
	"sort"
)

func arrayExample() {
	// 数组主要有三种声明方式
	var arr [3]int
	var arr2 = [4]int{1, 2, 3, 4}
	arr3 := [...]int{2, 3, 4} // 语法糖，依靠编译器进行数组长度的推断。

	fmt.Printf("arr: %+v, len: %d\n", arr, len(arr))
	fmt.Printf("arr2: %+v, len: %d\n", arr2, len(arr2))
	fmt.Printf("arr3: %+v, len: %d\n", arr3, len(arr3))

	// 数组值复制
	arr4 := arr2
	arr4[0] = 4
	fmt.Printf("arr4: %+v, arr2: %+v\n", arr4, arr2)
}

func sliceExample() {
	// 声明切片
	var slice1 []int
	slice2 := make([]int, 5)
	slice3 := make([]int, 0, 5)
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Printf("slice1: %+v, len: %d, cap: %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %+v, len: %d, cap: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3: %+v, len: %d, cap: %d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("numbers: %+v, len: %d, cap: %d\n", numbers, len(numbers), cap(numbers))

	// 切片的截取
	num1 := numbers[1:3] // 截取范围：[1,3)
	num2 := numbers[:2]  // [0,2)
	num3 := numbers[2:]  // [2,len(num3))
	fmt.Printf("num1: %+v, len: %d, cap: %d\n", num1, len(num1), cap(num1))
	fmt.Printf("num2: %+v, len: %d, cap: %d\n", num2, len(num2), cap(num2))
	fmt.Printf("num3: %+v, len: %d, cap: %d\n", num3, len(num3), cap(num3))

	// 切片的复制，是底层指针的值复制，可以理解为数据的引用传递
	num4 := numbers
	num4[0] = 4
	fmt.Printf("num4: %+v, numbers: %+v\n", num4, numbers)

	// 切片添加元素
	num5 := make([]int, 0)
	num5 = append(num5, 0)
	num5 = append(num5, 1, 2, 3)
	num5 = append(num5, []int{4, 5, 6}...)
	fmt.Printf("num5: %+v\n", num5)

	// 切片删除元素
	num5 = num5[1:] // 删除第一个元素
	fmt.Printf("num5: %+v\n", num5)
	num5 = num5[:len(num5)-1] // 删除最后一个元素
	fmt.Printf("num5: %+v\n", num5)
	num5 = append(num5[:2], num5[3:]...) // 删除下标为2的元素
	fmt.Printf("num5: %+v\n", num5)

	// 切片的深拷贝，新切片的修改不会影响旧切片
	num6 := []int{1, 2, 3}
	num7 := make([]int, len(num6), len(num6))
	copy(num7, num6)
	num7[0] = 4
	fmt.Printf("num6: %+v, len: %d, cap: %d\n", num6, len(num6), cap(num6))
	fmt.Printf("num7: %+v, len: %d, cap: %d\n", num7, len(num7), cap(num7))
}

/*
[167. 两数之和 II - 输入有序数组](https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/)

题目大意：输入一个递增排序的数组和一个值k，请问如何在数组中找出两个和为k的数字并返回它们的下标？
假设数组中存在且只存在一对符合条件的数字，同时一个数字不能使用两次。下标从1开始。
例如，输入数组[1，2，4，6，10]，k的值为8，数组中的数字2与6的和为8，由于下标从1开始，则它们的下标分别为2与4。
*/
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		} else if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

/*
[15. 三数之和](https://leetcode-cn.com/problems/3sum/)

题目大意：输入一个数组，如何找出数组中所有和为0的3个数字的三元组？需要注意的是，返回值中不得包含重复的三元组。
例如，在数组[-1，0，1，2，-1，-4]中有两个三元组的和为0，它们分别是[-1，0，1]和[-1，-1，2]。
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	last := len(nums) - 1
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, last

		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})

				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
				j++
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}

		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

/*
[209. 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

题目大意：输入一个正整数组成的数组和一个正整数k，请问数组中和大于或等于k的连续子数组的最短长度是多少？如果不存在所有数字之和大于或等于k的子数组，则返回0。
例如，输入数组[5，1，4，3]，k的值为7，和大于或等于7的最短连续子数组是[4，3]，因此输出它的长度2。
*/
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	res := math.MaxInt
	li := 0
	for ri := 0; ri < len(nums); ri++ {
		sum += nums[ri]
		for sum >= target {
			res = min(res, ri-li+1)
			if res == 1 {
				return res
			}
			sum -= nums[li]
			li++
		}
	}

	if res == math.MaxInt {
		return 0
	}
	return res
}

/*
[713. 乘积小于K的子数组](https://leetcode-cn.com/problems/subarray-product-less-than-k/)

题目大意：输入一个由正整数组成的数组和一个正整数k，请问数组中有多少个数字乘积小于k的连续子数组？
例如，输入数组[10，5，2，6]，k的值为100，有8个子数组的所有数字的乘积小于100，它们分别是[10]、[5]、[2]、[6]、[10，5]、[5，2]、[2，6]和[5，2，6]。
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	res := 0
	product := 1

	li := 0
	for ri := 0; ri < len(nums); ri++ {
		product *= nums[ri]
		for product >= k && li <= ri {
			product /= nums[li]
			li++
		}
		res += ri - li + 1
	}
	return res
}

/*
[560. 和为 K 的子数组](https://leetcode-cn.com/problems/subarray-sum-equals-k/)

题目大意：输入一个整数数组和一个整数k，请问数组中有多少个数字之和等于k的连续子数组？
例如，输入数组[1，1，1]，k的值为2，有2个连续子数组之和等于2。
*/
func subarraySum(nums []int, k int) int {
	cnt := make(map[int]int)
	count := 0
	sum := 0

	for _, num := range nums {
		sum += num
		if sum == k {
			count++
		}
		count += cnt[sum-k]
		cnt[sum]++
	}
	return count
}

/*
[525. 连续数组](https://leetcode-cn.com/problems/contiguous-array/)

题目大意：输入一个只包含0和1的数组，请问如何求0和1的个数相同的最长连续子数组的长度？
例如，在数组[0，1，0]中有两个子数组包含相同个数的0和1，分别是[0，1]和[1，0]，它们的长度都是2，因此输出2。
*/
func findMaxLength(nums []int) int {
	cnt := make(map[int]int)
	maxLen := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		if sum == 0 {
			maxLen = max(i+1, maxLen)
		}

		if v, ok := cnt[sum]; ok {
			maxLen = max(i-v, maxLen)
		} else {
			cnt[sum] = i
		}
	}
	return maxLen
}

/*
[724. 寻找数组的中心下标](https://leetcode-cn.com/problems/find-pivot-index/)

题目大意：输入一个整数数组，如果一个数字左边的子数组的数字之和等于右边的子数组的数字之和，那么返回该数字的下标。如果存在多个这样的数字，则返回最左边一个数字的下标。如果不存在这样的数字，则返回-1。
例如，在数组[1，7，3，6，2，9]中，下标为3的数字（值为6）的左边3个数字1、7、3的和与右边两个数字2和9的和相等，都是11，因此正确的输出值是3。
*/
func pivotIndex(nums []int) int {
	left := 0
	right := 0
	for i := 0; i < len(nums); i++ {
		right += nums[i]
	}

	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}

/*
[304. 二维区域和检索 - 矩阵不可变](https://leetcode-cn.com/problems/range-sum-query-2d-immutable/)

题目大意：输入一个二维矩阵，如何计算给定左上角坐标和右下角坐标的子矩阵的数字之和？对于同一个二维矩阵，计算子矩阵的数字之和的函数可能由于输入不同的坐标而被反复调用多次。
例如，输入图中的二维矩阵，以及左上角坐标为（2，1）和右下角坐标为（4，3）的子矩阵，该函数输出8。
*/
type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sum := make([][]int, len(matrix)+1)
	sum[0] = make([]int, len(matrix[0])+1)
	for i := 1; i < len(sum); i++ {
		row := make([]int, len(matrix[0])+1)
		rowSum := 0
		for j := 1; j < len(row); j++ {
			rowSum += matrix[i-1][j-1]
			row[j] = rowSum + sum[i-1][j]
		}
		sum[i] = row
	}
	return NumMatrix{
		sum: sum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row1][col2+1] - this.sum[row2+1][col1] + this.sum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
