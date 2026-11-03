package markdown

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func MarkToHTML(mk []byte) string {
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	markdownData := mk

	var buf bytes.Buffer

	// Converte o Markdown para HTML
	if err := md.Convert(markdownData, &buf); err != nil {
		panic(err)
	}

	return buf.String()
}

func RemoveMetadata(data string) string {
	reg := regexp.MustCompile(`(?mi)---[\s\S]+---`)

	return strings.TrimLeft(reg.ReplaceAllString(data, ""), " ")
}

// Deprecated methods
func AddJustifyTag(data string) string {
	strs := strings.Split(strings.TrimSpace(data), "\n")
	formmated := ""
	left := 0

	for left < len(strs) {
		line := strs[left]
		if isEmpty(line) {
			// increment because current left is an empty character
			left += 1
			formmated += mountPara(strs, left)
		} else {
			formmated += line + "\n"
		}

		left += 1
	}

	return formmated
}

func mountPara(strs []string, left int) string {
	if !validateLine(strs, left) {
		return "\n" + strs[left] + "\n"
	}

	// inside p tag the markdown don't work
	withTags := addHtmlTags(strs[left])

	return `
<p align="justify">
` + withTags + "\n</p>\n"
}

func validateLine(strs []string, idx int) bool {
	newI := idx
	nextLine := newI + 1

	// Is end of file or next line is a \n
	isNextLineValid := nextLine >= len(strs) || strs[nextLine] == ""

	return len(strings.TrimSpace(strs[newI])) > 100 && isNextLineValid
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func addHtmlTags(str string) string {
	regLink := regexp.MustCompile(`(?mi)\[([^\]]+)\]\(([^\)]+)\)`)
	regAsterisk := regexp.MustCompile(`(?mi)\*{1,2}([^\*]+)\*{1,2}`)
	repl := regAsterisk.ReplaceAllStringFunc(str, func(s string) string {
		if strings.Count(s, "*") == 4 {
			return "<b>" + strings.ReplaceAll(s, "*", "") + "</b>"
		}

		return "<i>" + strings.ReplaceAll(s, "*", "") + "</i>"
	})
	repl = regLink.ReplaceAllString(repl, "<a href=\"$2\">$1</a>")

	return repl
}
