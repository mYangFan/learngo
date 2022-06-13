package parser

import (
	"gonb/crawler/engine"
	"gonb/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>(\d+)</td>`)
var nameRe = regexp.MustCompile(`<th><a href="http://album.zhenai.com/u/[0-9]+"[^>]*>([^<]+)</a></th>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(exactString(contents, ageRe))

	if err == nil {
		profile.Age = age
	}

	profile.Name = exactString(contents, nameRe)

	result := engine.ParseResult{
		Items: []interface{}{
			profile,
		},
	}

	return result
}

func exactString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}else{
		return ""
	}
}
