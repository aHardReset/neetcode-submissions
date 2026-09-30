func characterReplacement(s string, k int) int {
	l, r := 0, 0
	res := 0
	memory := [26]int{}

	for r < len(s) {
		localLength, isEmpty := localCharacterReplacement(memory, k)
		res = max(res, localLength)
		if localLength == 0 && !isEmpty{
			memory[s[l]-'A']--
			l++
		} else{
			memory[s[r]-'A']++
			r++
		}
	}
	localLength, _ := localCharacterReplacement(memory, k)
	res = max(res, localLength)
	return res

}

func localCharacterReplacement(memory [26]int, k int) (int, bool){
	most := getMaxFromMemory(memory)
	if most == 0{
		return 0, true
	}
	flag := false
	changes := 0
	length := 0
	for _, m := range memory {
		length = length + m
		if m == most && flag == false {
			flag = true
			continue
		}
		if m < most || flag == true {
			changes = changes + m
		}
	}
	if changes <= k {
		return length, false
	} else {
		return 0, false
	}
}

func getMaxFromMemory(memory [26]int) int {
	curr := memory[0]
	for _, v := range memory {
		curr = max(curr, v)
	}
	return curr
}
// 2
// [1,1,1]

// ABC