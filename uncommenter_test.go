package jsonuncommenter

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestRemoveComment(t *testing.T) {

	js := struct{ in, out string }{
		`
{
    "seeds": [
        "http://t66y.com/thread0806.php?fid=16"// http下载
//
      // "http://t66y.com/thread0806.php?fid=21"
    ]
}`, `
{
    "seeds": [
        "http://t66y.com/thread0806.php?fid=16"
    ]
}`}

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
