package goraph

import (
	"regexp"
	"strings"
	"io/ioutil"
)

type Graph struct {
	root GraphNode
}

type GraphNode struct {
	children []GraphNode
	label string
}

func (g *GraphNode) AddChild(level int, label string) {
	var parent = g
	for level > 1 {
		level--;
		parent = &parent.children[len(parent.children)-1]
	}
	var child GraphNode
	child.label = label
	parent.children = append(parent.children, child)
}

type Translater interface {
	Convert(graph *GraphNode) string
}

var nodeRegexp = regexp.MustCompile(`^(-+)>\s*(.+)`)

func stringToGraph(src string) (*GraphNode, error) {
	var graph GraphNode
	var initialized = false
	for _, l := range strings.Split(src, "\n") {
		if !nodeRegexp.MatchString(l) {
			continue
		}
		var dashes = nodeRegexp.ReplaceAllString(l, "$1")
		var level = len(dashes) - 1
		var label = nodeRegexp.ReplaceAllString(l, "$2")
		if !initialized {
			if level == 0 {
				graph.label = label
				initialized = true
			}
		} else {
			graph.AddChild(level, label)
		}
	}

	return &graph, nil
}

func ConvertString(src string, trans Translater) (string, error) {
	var graph, err = stringToGraph(src)
	if err != nil {
		return "", err
	}

	output := trans.Convert(graph)

	return output, nil
}

func ConvertFile(path string, trans Translater) (string, error) {
	var buf, err = ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return ConvertString(string(buf), trans)
}
