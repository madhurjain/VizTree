package autocomplete

import "testing"

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestSearchEmpty(t *testing.T) {
	trie := NewTrie()
	words := trie.SearchWords("a")
	assert(t, len(words) == 0)

	trie.AddWord("apple")
	words = trie.SearchWords("")
	assert(t, len(words) == 0)
}

func TestSearchSingle(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("apple")
	trie.AddWord("pie")
	words := trie.SearchWords("p")
	assert(t, len(words) == 1)
	assert(t, words[0] == "pie")
}

func TestSearchMultiple(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("apple")
	trie.AddWord("pie")
	trie.AddWord("cake")
	trie.AddWord("pastry")
	trie.AddWord("coal")
	trie.AddWord("coke")
	trie.AddWord("coconut")
	trie.AddWord("confectionery")
	words := trie.SearchWords("co")

	assert(t, len(words) == 4)
	assert(t, words[0] == "coal" && words[1] == "coke" && words[2] == "coconut" && words[3] == "confectionery")
}

func TestSearchCaseSensitive(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("Ape")
	trie.AddWord("Apple")
	trie.AddWord("apricot")
	trie.AddWord("Ant")
	trie.AddWord("Ace")
	trie.AddWord("act")
	trie.AddWord("action")
	words := trie.SearchWords("Ap")

	assert(t, len(words) == 2)
	assert(t, words[0] == "Ape" && words[1] == "Apple")
}

func TestSearchUnicode(t *testing.T) {
	trie := NewTrie()
	trie.AddWord("おはよう")
	trie.AddWord("こんにちは")
	trie.AddWord("こんばんは")
	trie.AddWord("おやすみ")
	trie.AddWord("ありがとう")

	words := trie.SearchWords("こん")
	assert(t, len(words) == 2)
}
