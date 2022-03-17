package sword

// 3.1 字符串的基础知识

// 3.2 双指针

/*
https://leetcode-cn.com/problems/MPnaiL/

面试题14：字符串中的变位词

题目：输入字符串s1和s2，如何判断字符串s2中是否包含字符串s1的某个变位词？如果字符串s2中包含字符串s1的某个变位词，则字符串s1至少有一个变位词是字符串s2的子字符串。假设两个字符串中只包含英文小写字母。
例如，字符串s1为"ac"，字符串s2为"dgcaf"，由于字符串s2中包含字符串s1的变位词"ca"，因此输出为true。如果字符串s1为"ab"，字符串s2为"dgcaf"，则输出为false。
*/
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	a1 := [26]int{}
	a2 := [26]int{}
	for i := 0; i < len(s1); i++ {
		a1[s1[i]-'a']++
		a2[s2[i]-'a']++
	}
	if a1 == a2 {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		a2[s2[i]-'a']++
		a2[s2[i-len(s1)]-'a']--
		if a2 == a1 {
			return true
		}
	}
	return false
}

/*
https://leetcode-cn.com/problems/VabMRr/

面试题15：字符串中的所有变位词

题目：输入字符串s1和s2，如何找出字符串s2的所有变位词在字符串s1中的起始下标？假设两个字符串中只包含英文小写字母。
例如，字符串s1为"cbadabacg"，字符串s2为"abc"，字符串s2的两个变位词"cba"和"bac"是字符串s1中的子字符串，输出它们在字符串s1中的起始下标0和5。
*/
func findAnagrams(s string, p string) []int {
	res := make([]int, 0)

	if len(s) < len(p) {
		return res
	}

	a1 := [26]int{}
	a2 := [26]int{}
	for i := 0; i < len(p); i++ {
		a1[p[i]-'a']++
		a2[s[i]-'a']++
	}
	if a1 == a2 {
		res = append(res, 0)
	}
	for i := 0; i < len(s)-len(p); i++ {
		a2[s[i+len(p)]-'a']++
		a2[s[i]-'a']--
		if a2 == a1 {
			res = append(res, i+1)
		}
	}
	return res
}
