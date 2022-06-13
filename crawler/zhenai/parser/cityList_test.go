package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	const resultSize = 470

	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+"request; but had %d", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+"item; but had %d", resultSize, len(result.Items))
	}

	for i, city := range expectedCities {
		if result.Items[i] != city {
			t.Errorf("expected city #%d: %s; but "+"was %s", i, city, result.Items[i])
		}
	}
}
