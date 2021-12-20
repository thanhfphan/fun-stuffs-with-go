package main

import (
	"flag"
	"fmt"

	"github.com/goccy/go-graphviz/cgraph"
	"github.com/thanhfphan/fun-stuffs-with-go/data"
	"github.com/thanhfphan/fun-stuffs-with-go/trie"
)

func main() {
	t := trie.NewPrefix()

	for _, name := range data.Names {
		t.Add(name)
	}
	fmt.Println("build data succeed *")

	var userSearch string
	flag.StringVar(&userSearch, "name", "name", "input the name")
	flag.Parse()

	fmt.Println(fmt.Sprintf("user enter %s", userSearch)) //nolint

	result := t.AutoComplete(userSearch)
	fmt.Println(len(result))
	for _, r := range result {
		fmt.Println(r)
	}
	//Render
	// g := graphviz.New()
	// graph, err := g.Graph()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func() {
	// 	graph.Close()
	// 	g.Close()
	// }()

	// travelTrie(graph, t.Root)

	// if err := g.RenderFilename(graph, graphviz.PNG, "./graph.png"); err != nil {
	// 	log.Fatal(err)
	// }
}

func travelTrie(g *cgraph.Graph, node *trie.Node) {
	if node == nil {
		return
	}
	nStart, _ := g.CreateNode(node.Value)
	for _, c := range node.Childs {
		nEnd, _ := g.CreateNode(c.Value)
		e, _ := g.CreateEdge(node.Value, nStart, nEnd)
		e.SetLabel(node.Value)
		travelTrie(g, c)
	}
}
