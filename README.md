# Json uncommenter

remove comment from the json file.

1. Only line comment `//` is supported.
2. Multi-line comment `/**/` is not supported.

## Example

please find the example in [example folder](./example)

main.go

```go
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

```

config.json

```json
{
  "seeds": [
    "https://www.baidu.com" // http下载
    //
    // "https://v.qq.com"
  ],
  "name": "test"
}
```

output

```sh
$ go run main.go
{Seeds:[http://t66y.com/thread0806.php?fid=16] Name:test}
```

reference

- [strip-out-c-style-comments-from-a-byte]( https://stackoverflow.com/questions/12682405/strip-out-c-style-comments-from-a-byte)
