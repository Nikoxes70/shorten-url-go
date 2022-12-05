package parser

import (
	"fmt"
	"strconv"
	"strings"

	"shorten-url-go/shorturl/shorturl"
)

type parser struct {
}

func NewParser() *parser {
	return &parser{}
}

func (p parser) RawToRedirections(raw map[string]interface{}) ([]shorturl.Redirection, error) {
	var redirections []shorturl.Redirection
	for key, rawURL := range raw {
		from, to, err := p.keyToInterval(key)
		if err != nil {
			return nil, err
		}
		redirection := shorturl.Redirection{
			From: from,
			To:   to,
			URL:  p.rawToString(rawURL),
		}
		redirections = append(redirections, redirection)
	}
	return redirections, nil
}

func (p parser) keyToInterval(key string) (int64, int64, error) {
	interval := strings.Split(key, "-")
	if len(interval) != 2 {
		return 0, 0, fmt.Errorf("wrong key interval")
	}

	from, err := p.stringToInt64(interval[0])
	if err != nil {
		return 0, 0, fmt.Errorf("wrong from value")
	}

	to, err := p.stringToInt64(interval[1])
	if len(interval) != 2 {
		return 0, 0, fmt.Errorf("wrong to value")
	}
	return from, to, nil
}

func (p parser) rawToString(raw interface{}) string {
	if raw == nil {
		return ""
	}

	if v, ok := raw.(string); ok {
		return v
	}
	return string(raw.([]byte)[:])
}

func (p parser) stringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
