package entities

import (
	"crypto/md5"
	"fmt"
	"io"
)

type UrlPath struct {
	token string
	Path  string
}

func (url *UrlPath) WithToken(token string) {
	url.token = token
}

func (url *UrlPath) String() string {
	return fmt.Sprintf("%s%s", url.token, url.Path)
}

func (url *UrlPath) GetMD5Path() string {
	hash := md5.New()
	io.WriteString(hash, url.String())

	return fmt.Sprintf("%x", hash.Sum(nil))
}
