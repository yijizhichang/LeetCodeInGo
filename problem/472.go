package main

type node struct {
	child [26]*node
	empty bool
}

func NewNode() *node {
	return &node{
		child: [26]*node{},
		empty: true,
	}
}

func findAllConcatenatedWordsInADict(words []string) []string {
	root := NewNode()
	for _, word := range words {
		curr := root
		for _, v := range word {
			index := v - 97
			if curr.child[index] == nil {
				curr.child[index] = NewNode()
			}
			curr = curr.child[index]
		}
		curr.empty = false
	}
	var res []string
	for _, word := range words {
		if Contains(word, root, 0) {
			res = append(res, word)
		}
	}

	return res
}

func Contains(word string, root *node, count int) bool {
	curr := root
	for k, v := range word {
		curr = curr.child[v-97]
		if curr == nil {
			return false
		}
		if !curr.empty {
			if k >= len(word)-1 {
				return count > 0
			}
			// cat at a
			ww := word[k+1:]
			if Contains(string(ww), root, count+1) {
				return true
			}
		}
	}
	return false
}

