package jsonuncommenter

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestRemoveComment(t *testing.T) {

	cases := []struct{ in, out string }{{
		`
{
    "seeds": [
        "https://www.baidu.com"// http下载
//
      // "https://v.qq.com"
    ]
}`, `
{
    "seeds": [
        "https://www.baidu.com"
    ]
}`}}

	for _, js := range cases {
		reader := strings.NewReader(js.in)
		newReader := RemoveComment(reader)

		bs, err := ioutil.ReadAll(newReader)
		if err != nil {
			log.Fatal(err)
		}

		out := string(bs)

		if out != js.out {
			t.Fatal("want:", js.out, "got: ", out)
		}
	}
}
