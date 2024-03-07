package HashTable

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
//示例 2: 输入: s = "rat", t = "car" 输出: false
//说明: 你可以假设字符串只包含小写字母。

//func isAnagram(s string, t string) bool {
//	record := [26]int{}
//	for _, r := range s {
//
//	}
//}

// 只对字符串遍历一次
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) { // 长短不一
		return false
	}
	records := [26]int{}
	for index := 0; index < len(s); index++ {
		if s[index] == t[index] {
			continue //对相同的元素跳过
		}
		//同时遍历两个字符串
		sCharIndex := s[index] - 'a' //下标索引
		records[sCharIndex]++
		tCharIndex := t[index] - 'a'
		records[tCharIndex]--
	}
	for _, record := range records {
		if record != 0 {
			return false
		}
	}
	return true
}
