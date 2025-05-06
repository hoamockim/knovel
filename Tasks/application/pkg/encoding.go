package pkg

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func EncodeJsonFile(filePath string, rs interface{}) error {
	var des io.Writer = os.Stdout
	if strings.TrimSpace(filePath) == "" {
		return errors.New("this file is not exist")
	}
	if _, err := os.Open(filePath); os.IsExist(err) {
		// generate new file path
		filePath = fmt.Sprintf("%v_%v", filePath, UUID8())
		return EncodeJsonFile(filePath, rs)
	}
	return json.NewEncoder(des).Encode(rs)
}

func DecodeJsonFile(filePath string, rs interface{}) error {
	var src io.Reader = os.Stdin
	if strings.TrimSpace(filePath) == "" {
		return errors.New("this file is not exist")
	}

	if f, err := os.Open(filePath); err != nil {
		return err
	} else {
		defer f.Close()
		src = f
	}

	return json.NewDecoder(src).Decode(rs)
}

func EncodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func DecodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}
