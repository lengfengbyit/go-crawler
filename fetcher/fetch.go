package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil{
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error status code: %d", response.StatusCode)
	}

	// 转换编码
	bodyReader := bufio.NewReader(response.Body)
	encode := GetEncoding(bodyReader)
	reader := transform.NewReader(bodyReader, encode.NewDecoder())

	return ioutil.ReadAll(reader)
}

// 获取编码
func GetEncoding(r *bufio.Reader) encoding.Encoding {

	// 读取1024个字节
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("fetch error: %v", err)
		return unicode.UTF8
	}

	// 获取编码
	encode, _, _ := charset.DetermineEncoding(bytes, "")
	return encode
}
