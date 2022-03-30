package datastructure

type TrieNode[T Trier] struct {
	Data     T
	Children []*TrieNode[T]
	NodeType
	TopN TopN[T]
}

func NewTrieNode[T Trier](t T) *TrieNode[T] {
	trieNode := TrieNode[T]{Data: t, Children: make([]*TrieNode[T], 0), NodeType: Keyword}
	trieNode.TopN = MakeTopN1(&trieNode.Data)
	return &trieNode
}

func NewTrieNode1[T Trier]() *TrieNode[T] {
	return &TrieNode[T]{Children: make([]*TrieNode[T], 0), NodeType: Sentinel}
}

func NewTrieNode2[T Trier](t T, nodeType NodeType) *TrieNode[T] {
	trieNode := TrieNode[T]{Data: t, Children: make([]*TrieNode[T], 0), NodeType: nodeType}
	if nodeType == Keyword {
		trieNode.TopN = MakeTopN1(&trieNode.Data)
	} else {
		trieNode.TopN = MakeTopN[T]()
	}
	return &trieNode
}

func (node *TrieNode[T]) Add(t T) (bool, error) {
	defer func() {
		if node.NodeType == Keyword {
			node.TopN = MakeTopN1(&node.Data)
		} else {
			node.TopN = MakeTopN[T]()
		}
		for _, c := range node.Children {
			node.TopN.Merge(c.TopN)
		}
	}()

	for _, c := range node.Children {
		childCompareResult := t.Compare(c.Data)
		if childCompareResult == Equal {
			newData, _ := c.Data.Add(t)
			c.Data = newData.(T)
			if c.NodeType == Keyword {
				c.TopN.Sort()
			} else {
				c.NodeType = Keyword
				c.TopN.Merge(MakeTopN1(&node.Data))
			}
			return true, nil
		}
		if childCompareResult == IsChild || childCompareResult == IsDescendent {
			return c.Add(t)
		}
	}

	newData, err := t.GetChildOf(node.Data)
	if err != nil {
		return false, err
	}
	if newData.Compare(t) == Equal {
		node.Children = append(node.Children, NewTrieNode(t))
		return true, nil
	}
	newAncestor := NewTrieNode2(newData.(T), Intermediate)
	node.Children = append(node.Children, newAncestor)
	return newAncestor.Add(t)
}

func (node *TrieNode[T]) Find(t T) *TrieNode[T] {

	compareResult := node.Data.Compare(t)
	if compareResult == Equal {
		return node
	}

	if compareResult == IsParent || compareResult == IsAncestor {
		for _, c := range node.Children {
			if target := c.Find(t); target != nil {
				return target
			}
		}
	}
	return nil
}

type NodeType int64

const (
	Sentinel NodeType = iota
	Intermediate
	Keyword
)
