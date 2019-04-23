package jsonuncommenter

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

func RemoveComment(reader io.Reader) (newReader io.Reader) {

	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	s := string(bs)
	//fmt.Println("before ", s)
	re1 := regexp.MustCompile(`(?im)^\s+\/\/.*$`) // 整行注释

	s = re1.ReplaceAllString(s, "")
	//fmt.Println("after1 ", s)
	re2 := regexp.MustCompile(`(?im)\/\/[^"\[\]]+$`) // 行末
	s = re2.ReplaceAllString(s, "")
	//fmt.Println("after2 ", s)
	newReader = strings.NewReader(s)
	return

}
