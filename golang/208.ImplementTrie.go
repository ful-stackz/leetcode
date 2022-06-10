package leetcode

type Trie struct {
	children []*Trie
	isWord   bool
}

func Constructor() Trie {
	return Trie{children: make([]*Trie, 26), isWord: false}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		symbol := word[i] - 'a'

		if node.children[symbol] == nil {
			node.children[symbol] = &Trie{children: make([]*Trie, 26), isWord: false}
		}

		node = node.children[symbol]
	}
	node.isWord = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		symbol := word[i] - 'a'

		if node.children[symbol] == nil {
			return false
		}

		node = node.children[symbol]
	}
	return node.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix) && node != nil; i++ {
		symbol := prefix[i] - 'a'
		node = node.children[symbol]
	}
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
