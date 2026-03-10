package trie

type Node struct {
	next map[rune]*Node
	end  bool
}

type Trie struct {
	head *Node
}

func NewTrie() *Trie {
	return &Trie{head: &Node{next: make(map[rune]*Node)}}
}

func (t *Trie) Add(word string) {
	if word == "" {
		return
	}

	curr := t.head
	for _, ch := range word {
		if curr.next[ch] == nil {
			curr.next[ch] = &Node{next: make(map[rune]*Node)}
		}
		curr = curr.next[ch]
	}
	curr.end = true
}

func (t *Trie) Exists(word string) bool {
	if word == "" {
		return false
	}

	curr := t.head
	for _, ch := range word {
		if curr.next[ch] == nil {
			return false
		}
		curr = curr.next[ch]
	}
	return curr.end
}

func (t *Trie) Delete(word string) {
	if word == "" {
		return
	}

	curr := t.head
	for _, ch := range word {
		if curr.next[ch] == nil {
			return
		}
		curr = curr.next[ch]
	}
	curr.end = false
}

func (t *Trie) List() []string {
	var result []string
	var dfs func(n *Node, prefix string)
	dfs = func(n *Node, prefix string) {
		if n.end {
			result = append(result, prefix)
		}
		for ch, child := range n.next {
			dfs(child, prefix+string(ch))
		}
	}
	dfs(t.head, "")
	return result
}
