package signature

import (
	"bytes"
	"net/http"
)

// The algorithm uses at most sniffLen bytes to make its decision.
const sniffLen = 512

// isWS reports whether the provided byte is a whitespace byte (0xWS)
// as defined in https://mimesniff.spec.whatwg.org/#terminology.
func isWS(b byte) bool {
	switch b {
	case '\t', '\n', '\x0c', '\r', ' ':
		return true
	}
	return false
}

// isTT reports whether the provided byte is a tag-terminating byte (0xTT)
// as defined in https://mimesniff.spec.whatwg.org/#terminology.
func isTT(b byte) bool {
	switch b {
	case ' ', '>':
		return true
	}
	return false
}

// SniffResult 识别结果
type SniffResult struct {
	// 某些签名对应多种文件格式，比如 zip,docx,xlsx，所以这里用数组~
	Extensions []string
	// 互联网媒体类型（Internet media type，也称为MIME类型（MIME type）或内容类型（content type））
	// 是给互联网上传输的内容赋予的分类类型。
	ContentType string
}

var voidSniffResult SniffResult = SniffResult{}

type sniffSig interface {
	// match returns the MIME type of the data, or "" if unknown.
	match(data []byte, firstNonWS int) SniffResult
}

var sniffSignatures = make([]sniffSig, 0)

var mapContentType = make(map[string]SniffResult)

func Sniff(data []byte) SniffResult {
	// 使用 Go 自带识别出 ContentType 然后查看是否注册了对应的 SniffResult
	ct := http.DetectContentType(data)
	if sr, ok := mapContentType[ct]; ok {
		return sr
	}

	if len(data) > sniffLen {
		data = data[:sniffLen]
	}

	// Index of the first non-whitespace byte in data.
	firstNonWS := 0
	for ; firstNonWS < len(data) && isWS(data[firstNonWS]); firstNonWS++ {
	}

	for _, sig := range sniffSignatures {
		if sr := sig.match(data, firstNonWS); len(sr.Extensions) != 0 {
			sr.ContentType = ct
			return sr
		}
	}

	return voidSniffResult
}

type contentSig struct {
	// ContentType
	ct string
	sr SniffResult
}

func (c *contentSig) match(data []byte, firstNonWS int) SniffResult {
	ct := http.DetectContentType(data)
	if ct == c.ct {
		return c.sr
	}
	return voidSniffResult
}

type exactSig struct {
	// file signatures
	sig []byte
	//
	sr SniffResult
}

func (e *exactSig) match(data []byte, firstNonWS int) SniffResult {
	if bytes.HasPrefix(data, e.sig) {
		return e.sr
	}
	return voidSniffResult
}

type maskedSig struct {
	mask, pat []byte
	skipWS    bool
	ct        string
}

func (m *maskedSig) match(data []byte, firstNonWS int) SniffResult {
	// pattern matching algorithm section 6
	// https://mimesniff.spec.whatwg.org/#pattern-matching-algorithm

	if m.skipWS {
		data = data[firstNonWS:]
	}
	if len(m.pat) != len(m.mask) {
		return voidSniffResult
	}
	if len(data) < len(m.pat) {
		return voidSniffResult
	}
	for i, pb := range m.pat {
		maskedData := data[i] & m.mask[i]
		if maskedData != pb {
			return voidSniffResult
		}
	}
	return voidSniffResult
}
