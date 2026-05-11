package util

import (
	"github.com/gomarkdown/markdown"
	"strings"
)

func FormatAttr(str string) string {

	tempSlice := strings.Split(str, "\n")
	var tempStr string
	for _, v := range tempSlice {
		md := []byte(v)
		output := markdown.ToHTML(md, nil, nil)
		tempStr += string(output)
	}
	return tempStr
}
