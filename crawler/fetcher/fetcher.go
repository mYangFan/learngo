package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"gonb/pkg/log"
	"io/ioutil"
	"net/http"
	"time"
)


var ratelimiter = time.Tick(100 * time.Millisecond)
func Fetch(url string) ([]byte, error) {
	<-ratelimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("Wrong Status code:%d")
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	transform.NewReader(bodyReader, e.NewDecoder())
	return  ioutil.ReadAll(bodyReader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		//出错就返回utf8
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
