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

func headingTextWriter(w io.Writer, content []byte) {
	io.WriteString(w, fmt.Sprintf("[::b]%s\n", string(content)))
}
