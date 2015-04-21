VizTree
-------

VizTree is a Simple, Fast Trie based Auto Suggest

Demo - [http://viztree.madhur.me/](http://viztree.madhur.me/)

```
.
+-- server.go - Demo web application with D3 based visualization for tree
+-- linuxwords.txt - Sample word list
+-- autocomplete
    +-- autocomplete.go - Library for Trie based word search
```

Usage
-----

```go
package main

import "./autocomplete"
import "fmt"

func main() {
	tree := autocomplete.NewTrie()
	tree.AddWord("Ace")
	tree.AddWord("Ape")
	tree.AddWord("Apple")
	tree.AddWord("Banana")

	words := tree.SearchWords("Ap")
	fmt.Println(words) // Prints [Ape Apple]
}
```