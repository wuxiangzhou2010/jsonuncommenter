package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/wuxiangzhou2010/jsonuncommenter"
)

type seeds struct {
	Seeds []string `json:"seeds"`
	Name  string   `json:"name"`
}

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	f, err := os.Open(path.Join(cwd, `config.json`))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	newReader := jsonuncommenter.RemoveComment(f)

	jsonParser := json.NewDecoder(newReader)
	var s seeds
	jsonParser.Decode(&s)
	fmt.Printf("%+v", s)
}
