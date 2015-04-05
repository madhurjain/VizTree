package autocomplete

type Node struct {
	alpha   rune
	next    []*Node
	isFinal bool
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	trie := new(Trie)
	trie.root = &Node{alpha: 0, isFinal: false}
	trie.root.next = make([]*Node, 0)
	return trie
}

func (trie *Trie) AddWord(word string) {
	currentNode := trie.root
	// iterate on each character in the word
	for _, alpha := range word {
		var newNode *Node
		// check if a node exists for every character
		for _, nextNode := range currentNode.next {
			if alpha == nextNode.alpha {
				newNode = nextNode
				break
			}
		}
		// if a child node does not exist for char alpha, add it
		if newNode == nil {
			newNode = new(Node)
			newNode.alpha = alpha
			newNode.next = make([]*Node, 0)
			newNode.isFinal = false
			currentNode.next = append(currentNode.next, newNode)
		}
		currentNode = newNode
	}
	// ending node for the word
	currentNode.isFinal = true
}

func (trie *Trie) SearchWords(keyword string) []string {
	if keyword == "" {
		return []string{}
	}
	currentNode := trie.root
	var prefixFound bool
	// traverse for prefix
	for _, alpha := range keyword {
		prefixFound = false
		// iterate through tree for every char
		for _, child := range currentNode.next {
			if alpha == child.alpha {
				prefixFound = true
				currentNode = child
				// move to next char
				continue
			}
		}
	}
	if !prefixFound {
		return []string{}
	}
	// recursively traverse for available suffixes
	return trie.getWords(currentNode, keyword)
}

func (trie *Trie) getWords(node *Node, prefix string) (words []string) {
	for _, n := range node.next {
		// if node is final, add it to list of words
		if n.isFinal {
			word := prefix + string(n.alpha)
			words = append(words, word)
		}
		// recursively iterate for each child node
		words = append(words, trie.getWords(n, prefix+string(n.alpha))...)
	}
	return words
}
