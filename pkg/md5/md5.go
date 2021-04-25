package md5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func New(key string) string {
	hash := md5.New()
	io.WriteString(hash, key)
	return fmt.Sprintf("%x", hash.Sum([]byte{}))
}
