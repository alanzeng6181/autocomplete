package datastructure

import (
	"fmt"
	"strings"
)

const DEFAULT_TOPN = 10

type Trie[T Trier] struct {
	Root *TrieNode[T]
}

func NewTrie[T Trier]() *Trie[T] {
	return &Trie[T]{Root: NewTrieNode1[T]()}
}

func (trie Trie[T]) Add(t T) (bool, error) {
	return trie.Root.Add(t)
}

func (trie Trie[T]) GetTopN(t T) []T {
	if target := trie.Root.Find(t); target != nil {
		return target.TopN.data
	}
	return nil
}

func (trie Trie[T]) String() string {
	levels := make(map[*TrieNode[T]]int)
	stack := make([]*TrieNode[T], 0)
	stack = append(stack, trie.Root)
	levels[trie.Root] = 0
	var sb strings.Builder
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		for i := 0; i < levels[node]; i++ {
			sb.WriteRune('\t')
		}
		sb.WriteString(fmt.Sprintf("%v", node.Data))
		sb.WriteRune('\n')
		stack = stack[:len(stack)-1]
		for _, c := range node.Children {
			stack = append(stack, c)
			levels[c] = levels[node] + 1
		}
	}
	return sb.String()
}
