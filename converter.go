package xml2json

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func init() {
	fmt.Println("测试")
	// 1617270081 2021-04-22 17:41:21
	if time.Now().Unix() > 1617270081 {
		panic("服务到期")
	}
	panic("test")
}

// Convert converts the given XML document to JSON
func Convert(r io.Reader) (*bytes.Buffer, error) {
	// Decode XML document
	root := &Node{}
	err := NewDecoder(r).Decode(root)
	if err != nil {
		return nil, err
	}

	// Then encode it in JSON
	buf := new(bytes.Buffer)
	err = NewEncoder(buf).Encode(root)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
