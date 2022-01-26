package sword

import "fmt"

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
