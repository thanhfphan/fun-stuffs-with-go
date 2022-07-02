package main

import (
	"flag"
	"fmt"

	"github.com/thanhfphan/fun-stuffs-with-go/trie/data"
)

func main() {
	t := NewPrefix()

	for _, name := range data.Names {
		t.Add(name)
	}
	fmt.Println("build data succeed *")

	var userSearch string
	flag.StringVar(&userSearch, "name", "name", "input the name")
	flag.Parse()
	//FIXME: check err

	fmt.Sprintf("user enter %s\n", userSearch)

	result := t.AutoComplete(userSearch)
	fmt.Printf("found %d result\n", len(result))
	for _, r := range result {
		fmt.Println(r)
	}
}

type Node struct {
	Value  string
	Childs map[string]*Node
}

type PrefixTrie struct {
	Root *Node
}

func NewPrefix() *PrefixTrie {
	return &PrefixTrie{Root: &Node{Value: ""}}
}

func (t *PrefixTrie) Add(value string) {
	node := t.Root
	for _, c := range value {
		if node.Childs == nil {
			node.Childs = make(map[string]*Node)
		}

		n, ok := node.Childs[string(c)]
		if !ok {
			tmp := &Node{Value: string(c)}
			node.Childs[string(c)] = tmp
			node = tmp
			continue
		}

		node = n
	}

}

func (t *PrefixTrie) AutoComplete(input string) []string {
	var result []string
	node := t.Root

	for _, c := range input {
		if _, ok := node.Childs[string(c)]; !ok {
			return []string{}
		}

		node = node.Childs[string(c)]
	}

	for _, child := range node.Childs {
		result = append(result, walk(input, child)...)
	}

	return result
}

func walk(prefix string, t *Node) []string {
	if t == nil {
		return []string{prefix}
	}

	if t.Childs == nil || len(t.Childs) == 0 {
		return []string{prefix + t.Value}
	}

	var tmp []string
	for _, c := range t.Childs {
		tmp = append(tmp, walk(prefix+t.Value, c)...)
	}
	return tmp
}
