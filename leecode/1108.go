package main

var romanIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := len(s)
	total := 0
	for i := 0; i < n; i++ {
		value := romanIntMap[s[i]]
		if i < n-1 && value < romanIntMap[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}
