package trie

import (
	"testing"
)

func TestAdd(t *testing.T) {
	trie := NewPrefix()
	trie.Add("hello")
}
