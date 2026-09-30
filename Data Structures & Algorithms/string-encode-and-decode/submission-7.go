
type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    if len(strs) > 0 {
        res := strings.Join(strs, "[,]")
        if len(res) == 0 {
            return "[,]v"
        }
        return res
    } else {
        return ""
    }
    
}

func (s *Solution) Decode(str string) (res []string) {
    if str == "[,]v" {
        return []string{""}
    }
    if len(str) == 0 {
        return res
    }
    return strings.Split(str, "[,]")
}
