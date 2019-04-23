package jsonuncommenter

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

// RemoveComment remove comment from the io.Reader
func RemoveComment(reader io.Reader) (newReader io.Reader) {

	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	s := string(bs)
	re1 := regexp.MustCompile(`(?im)^\s+\/\/.*$`) // 整行注释
	s = re1.ReplaceAllString(s, "")
	re2 := regexp.MustCompile(`(?im)\/\/[^"\[\]]+$`) // 行末
	s = re2.ReplaceAllString(s, "")
	newReader = strings.NewReader(s)
	return
}
