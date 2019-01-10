package console

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/russross/blackfriday/v2"
)

func TestRender(t *testing.T) {
	testmdbytes, err := ioutil.ReadFile("testdata/test.md")
	if err != nil {
		t.Fatal(err)
	}
	renderer := new(Console)
	output := blackfriday.Run(testmdbytes, blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.CommonExtensions))
	fmt.Println(string(output))
}
