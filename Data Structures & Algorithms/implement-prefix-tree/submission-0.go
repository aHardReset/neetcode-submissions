
type PrefixTree struct {
	nodes map[rune]*PrefixTree
	isWord bool
}

func Constructor() PrefixTree {
    return PrefixTree{nodes: map[rune]*PrefixTree{}, isWord: false}
}

func (this *PrefixTree) Insert(word string) {
	current := this
	for _, c := range word {
		if _, ok := current.nodes[c]; !ok {
			current.nodes[c] = &PrefixTree{nodes: map[rune]*PrefixTree{}, isWord: false}
		}
		current = current.nodes[c]
	}
	current.isWord = true
}

func (this *PrefixTree) Search(word string) bool {
	current := this
	for _, c := range word {
		if _, ok := current.nodes[c]; !ok {
			return false
		}
		current = current.nodes[c]
	}
	return current.isWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	current := this
	for _, c := range prefix {
		if _, ok := current.nodes[c]; !ok {
			return false
		}
		current = current.nodes[c]
	}
	return true
}
