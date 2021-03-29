package xml2json

import (
	"bytes"
	"fmt"
	"io"
)

func init() {
	fmt.Println("矩阵科技")
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
