package goraph

type DotTranslater struct {
}

const dotTemplateBegin = "digraph {"
const dotTemplateEnd = "}"

func (d *DotTranslater) Convert(graph *GraphNode) string {
	var output = dotTemplateBegin + "\n" + d.convertNode("", graph) + dotTemplateEnd
	return output
}

func (d *DotTranslater) convertNode(cur string, node *GraphNode) string {
	var out = cur
	for _, child := range node.children {
		out += node.label + " -> " + child.label + "\n"
	}

	for _, child := range node.children {
		out = d.convertNode(out, &child)
	}

	return out
}
