package console

import (
	"fmt"
	"io"
	"strings"

	"github.com/russross/blackfriday"
)

func headingWriter(w io.Writer, heading blackfriday.HeadingData) {
	io.WriteString(w, fmt.Sprintf("\n%s\n", strings.Repeat("=", 20)))
}

func headingTextWriter(w io.Writer, content string) {
	io.WriteString(w, fmt.Sprintf("\n\n[::bu]%s\n\n", string(content)))
}

func codeWriter(w io.Writer, content string) {
	io.WriteString(w, fmt.Sprintf("\n\n"))
	for _, s := range strings.Split(content, "\n") {
		io.WriteString(w, fmt.Sprintf("[::d]%s[::-]\n", s))
	}
	io.WriteString(w, fmt.Sprintf("\n"))
}
