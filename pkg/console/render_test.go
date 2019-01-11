package console

import (
	"io/ioutil"
	"testing"

	"github.com/rivo/tview"
	"github.com/russross/blackfriday"
)

func TestRender(t *testing.T) {
	testmdbytes, err := ioutil.ReadFile("testdata/extensionbp.md")
	if err != nil {
		t.Fatal(err)
	}
	renderer := new(Console)
	output := blackfriday.Run(testmdbytes, blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.CommonExtensions))
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.Write(output)

	textView.SetBorder(true)
	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}

}
