package sword

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 二进制位运算示例
func binaryExample() {
	a := 60 // 60 = 00111100
	fmt.Printf("a = %d, %08b\n", a, a)
	b := 13 // 13 = 00001101
	fmt.Printf("b = %d, %08b\n", b, b)

	c := a & b // 12 = 00001100
	fmt.Printf("a & b = %d, %08b\n", c, c)

	c = a | b // 61 = 00111101
	fmt.Printf("a | b = %d, %08b\n", c, c)

	c = a ^ b // 49 = 00110001
	fmt.Printf("a ^ b = %d, %08b\n", c, c)

	c = a << 2 // 240 = 11110000
	fmt.Printf("a << 2 = %d, %08b\n", c, c)

	c = a >> 2 // 15 = 00001111
	fmt.Printf("a >> 2 = %d, %08b\n", c, c)
}

/*
[29. 两数相除](https://leetcode-cn.com/problems/divide-two-integers/)

题目：输入2个int型整数，它们进行除法计算并返回商，要求不得使用乘号'*'、除号'/'及求余符号'%'。当发生溢出时，返回最大的整数值。假设除数不为0。
例如，输入15和2，输出15/2的结果，即7。
*/
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	negative := 2
	if dividend > 0 {
		negative--
		dividend = -dividend
	}
	if divisor > 0 {
		negative--
		divisor = -divisor
	}

	result := 0
	halfMinInt32 := math.MinInt32 >> 1
	for dividend <= divisor {
		value := divisor
		quotient := 1
		for value >= halfMinInt32 && dividend <= value<<1 {
			quotient = quotient << 1
			value = value << 1
		}
		result += quotient
		dividend -= value
	}

	if negative == 1 {
		return -result
	}
	return result
}

/*
[67. 二进制求和](https://leetcode-cn.com/problems/add-binary/)

题目：输入两个表示二进制的字符串，请计算它们的和，并以二进制字符串的形式输出。
例如，输入的二进制字符串分别是"11"和"10"，则输出"101"。
*/
func addBinary(a string, b string) string {
	arr := make([]byte, 0)

	k := 0
	var carry uint8
	for k < len(a) || k < len(b) || carry > 0 {
		var r = carry
		if k < len(a) {
			r = r + a[len(a)-k-1] - '0'
		}
		if k < len(b) {
			r = r + b[len(b)-k-1] - '0'
		}
		k++

		if r == 2 {
			r = 0
			carry = 1
		} else if r == 3 {
			r = 1
			carry = 1
		} else {
			carry = 0
		}

		arr = append(arr, r+'0')
	}

	// 反转数组
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}

/*
[338. 比特位计数](https://leetcode-cn.com/problems/counting-bits/)

题目：输入一个非负数n，请计算0到n之间每个数字的二进制形式中1的个数，并输出一个数组。
例如，输入的n为4，由于0、1、2、3、4的二进制形式中1的个数分别为0、1、1、2、1，因此输出数组[0，1，1，2，1]。
*/
func countBits(n int) []int {
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i&(i-1)] + 1
	}
	return arr
}

func countBits2(n int) []int {
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i>>1] + (i & 1)
	}
	return arr
}

/*
[318. 最大单词长度乘积](https://leetcode-cn.com/problems/maximum-product-of-word-lengths/)

题目：输入一个字符串数组words，请计算不包含相同字符的两个字符串words[i]和words[j]的长度乘积的最大值。如果所有字符串都包含至少一个相同字符，那么返回0。假设字符串中只包含英文小写字母。
例如，输入的字符串数组words为["abcw"，"foo"，"bar"，"fxyz"，"abcdef"]，数组中的字符串"bar"与"foo"没有相同的字符，它们长度的乘积为9。
"abcw"与"fxyz"也没有相同的字符，它们长度的乘积为16，这是该数组不包含相同字符的一对字符串的长度乘积的最大值。
*/
func maxProduct(words []string) int {
	wordCnt := make([][26]bool, len(words))
	for i, word := range words {
		for _, ch := range word {
			wordCnt[i][ch-'a'] = true
		}
	}

	var res int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			k := 0
			for ; k < 26; k++ {
				if wordCnt[i][k] && wordCnt[j][k] {
					break
				}
			}

			if k == 26 {
				res = max(len(words[i])*len(words[j]), res)
			}
		}
	}
	return res
}

func maxProduct2(words []string) int {
	wordCnt := make([]int, len(words))
	for i, word := range words {
		for _, ch := range word {
			wordCnt[i] |= 1 << (ch - 'a')
		}
	}

	var res int
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if wordCnt[i]&wordCnt[j] == 0 {
				res = max(len(words[i])*len(words[j]), res)
			}
		}
	}
	return res
}
