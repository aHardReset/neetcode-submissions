import "slices"

func groupAnagrams(strs []string) (d [][]string) {
    groups := make(map[string][]string)
    
    for _, str := range strs {
        sortedStr := sortString(str)
        slice := groups[sortedStr]
        slice = append(slice, str)
        groups[sortedStr] = slice
    }

    for _, group := range groups {
        d = append(d, group)
    }

    return d
}

func sortString(str string) string {
    strSlice := make([]rune, len(str))

    for i, s := range str {
        strSlice[i] = s
    }

    slices.Sort(strSlice)

    res := ""
    for _, s := range strSlice {
        res += string(s)
    }
    return res
}