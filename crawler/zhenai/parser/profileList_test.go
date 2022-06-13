package parser

import (
	"gonb/crawler/fetcher"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")

	if err != nil {
		panic(contents)
	}

	ParseProfile(contents)
}