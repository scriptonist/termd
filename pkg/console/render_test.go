package console

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/hashnode/hashnode-cli/pkg/posts"
	"github.com/rivo/tview"
	"github.com/russross/blackfriday"
)

func TestRender(t *testing.T) {
	testmdbytes := []byte(getArticle(t, "cjqtwrdlt00g9ngs27ij20vob").Post.ContentMarkdown)
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

func getArticle(t *testing.T, id string) *posts.Post {
	var apiURL = fmt.Sprintf("https://hashnode.com/ajax/post/%s", id)
	client := http.Client{
		Timeout: time.Duration(1 * time.Minute),
	}
	resp, err := client.Get(apiURL)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	var body []byte
	if resp.StatusCode == http.StatusOK {
		// read response body
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
	}

	var post posts.Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		t.Fatal(err)
	}

	return &post
}
