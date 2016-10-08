package goraph

import (
	"testing"
)

func TestNodeRegexp(t *testing.T) {
	if !nodeRegexp.MatchString("-> lala") || !nodeRegexp.MatchString("--->Another test") || !nodeRegexp.MatchString("-->   Hehe") {
		t.Fail()
	}

	if nodeRegexp.MatchString("l-> foo") || nodeRegexp.MatchString("> bar") {
		t.Fail()
	}

	if nodeRegexp.ReplaceAllString("----> lala", "$1") != "----" || nodeRegexp.ReplaceAllString("-> foo", "$1") != "-" {
		t.Fail()
	}

	if nodeRegexp.ReplaceAllString("---> F00", "$2") != "F00" {
		t.Fail()
	}
}

func TestGraphNodeAddChild(t *testing.T) {
	var graph GraphNode
	graph.AddChild(1, "Hey here :)")
	graph.AddChild(1, "Another one !")
	graph.AddChild(2, "Foo")
	graph.AddChild(3, "Bar")
	if graph.children[0].label != "Hey here :)" ||
	graph.children[1].label != "Another one !" ||
	graph.children[1].children[0].label != "Foo" ||
	graph.children[1].children[0].children[0].label != "Bar" {
		t.Fail()
	}
}

func TestStringToGraph(t *testing.T) {
	var src = `-> One
--> Two
--> Three
---> Four`
	var graph, err = stringToGraph(src)
	if err != nil {
		t.Fail()
	}

	if graph.label != "One" ||
	graph.children[0].label != "Two" ||
	graph.children[1].label != "Three" ||
	graph.children[1].children[0].label != "Four" {
		t.Fail()
	}
}

const goodOutput = `digraph {
Root -> Foo
Root -> Other
Foo -> Bar
}`

func TestDotTranslater(t *testing.T) {
	var graph GraphNode
	var foo GraphNode
	var bar GraphNode
	var other GraphNode

	graph.label = "Root"
	other.label = "Other"
	foo.label = "Foo"
	bar.label = "Bar"

	foo.children = []GraphNode{bar}
	graph.children = []GraphNode{foo, other}

	var dot DotTranslater
	var out = dot.Convert(&graph)

	if out != goodOutput {
		t.Fail()
	}
}
