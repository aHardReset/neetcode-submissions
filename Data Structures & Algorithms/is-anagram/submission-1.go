func isAnagram(s string, t string) bool {
    if len(s) != len(s) {
        return false
    }

    sMap := make(map[rune]int)
    //tMap := make(map[rune]int)

    for _, v := range s {
        sMap[v]++
    }
    for _, v := range t{
        sMap[v]--
        if sMap[v] < 0{
            return false
        }
    }

    for _,v := range sMap {
        if v != 0 {
            return false
        }
    } 

    return true

}
