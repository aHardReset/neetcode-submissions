import "slices"

func groupAnagrams(strs []string) [][]string {
    var solution [][]string
    groups := make(map[string][]string)
    
    for _, str := range strs {
        sortedStr := sortString(str)
        groups[sortedStr] = append(groups[sortedStr], str)
    }

    for _, group := range groups {
        solution = append(solution, group)
    }

    return solution
}

func sortString(str string) string {
    strRunes := []rune(str)
    slices.Sort(strRunes)
    return string(strRunes)
}