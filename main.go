package main

import (
	"fmt"

	"github.com/alanzeng6181/autocomplete/datastructure"
)

func main() {
	trie := datastructure.NewTrie[datastructure.SearchString]()
	trie.Add(datastructure.MakeSearchString("abc", 100))
	trie.Add(datastructure.MakeSearchString("fabc", 50))
	trie.Add(datastructure.MakeSearchString("agbc", 100))
	trie.Add(datastructure.MakeSearchString("adbc", 100))
	trie.Add(datastructure.MakeSearchString("ggfabc", 1850))
	trie.Add(datastructure.MakeSearchString("zgzdfgfabc", 150))
	trie.Add(datastructure.MakeSearchString("adasdfbc", 100))
	trie.Add(datastructure.MakeSearchString("adbc", 100))
	trie.Add(datastructure.MakeSearchString("abdc", 60))
	trie.Add(datastructure.MakeSearchString("aadsfbc", 110))
	trie.Add(datastructure.MakeSearchString("tsdasdeaadsfbc", 110))
	fmt.Println(trie.String())
	fmt.Println(trie.GetTopN(datastructure.MakeSearchString1("abde")))
	fmt.Println(trie.GetTopN(datastructure.MakeSearchString1("zg")))
	fmt.Println(trie.GetTopN(datastructure.MakeSearchString1("zzz")))
}
