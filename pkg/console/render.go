package console

import (
	"io"

	blackfriday "github.com/russross/blackfriday/v2"
)

// Console implements blackfriday.Render
type Console struct {
}

// RenderNode is the main rendering method. It will be called once for
// every leaf node and twice for every non-leaf node (first with
// entering=true, then with entering=false). The method should write its
// rendition of the node to the supplied writer w.
func (c *Console) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) {

}

// RenderHeader  method will be passed an entire document tree, in case a particular
// implementation needs to inspect it to produce output.
//
// The output should be written to the supplied writer w. If your
// implementation has no header to write, supply an empty implementation.
func (c *Console) RenderHeader(w io.Writer, ast *blackfriday.Node) {

}

// RenderFooter is a symmetric counterpart of RenderHeader.
func RenderFooter(w io.Writer, ast *blackfriday.Node) {

}
