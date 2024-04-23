package model

import "bytes"

type Test struct {
	HttpMethod  string
	Path        string
	Body        *bytes.Buffer
	ContentType string
}
