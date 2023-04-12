package dataStructures

//PREFIX TREE
//Here trie representing lower cased letters, so each TreeNode can have at max 26 children. in the children array, a will be stored at zero index, b at 1 index etc..

type TrieTree struct {
	root *trieNode
}

type trieNode struct {
	children [26]*trieNode
	isWord   bool
}

func InitTrieTree() *TrieTree {
	return &TrieTree{
		&trieNode{},
	}
}

func (t *TrieTree) Insert(word string) {
	if t.root == nil {
		t.root = new(trieNode)
	}
	currentNode := t.root
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		wordIndex := word[i] - 'a'
		if currentNode.children[wordIndex] == nil {
			currentNode.children[wordIndex] = new(trieNode)
		}
		currentNode = currentNode.children[wordIndex]
	}
	currentNode.isWord = true
}

func (t *TrieTree) Find(word string) bool {
	if t.root == nil {
		return false
	}
	currentNode := t.root
	wordLen := len(word)
	for i := 0; i < wordLen; i++ {
		wordIndex := word[i] - 'a'
		if currentNode.children[wordIndex] == nil {
			return false
		}
		currentNode = currentNode.children[wordIndex]
	}
	return currentNode.isWord
}
