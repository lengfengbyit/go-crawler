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
	url2 "net/url"
)

func Fetch(url string) ([]byte, error) {
	urli := url2.URL{}
	urlProxy, _ := urli.Parse("http://xxx/")
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(urlProxy),
			MaxConnsPerHost: 20,
			MaxIdleConns: 20,
		},
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	if err != nil {
		return nil, err
	}

	request.Close = true
	//request.Header.Add("Connection", "close")
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


// 在解决问题之前需要了解关于go是如何实现connection的一些背景小知识：
//有两个协程，一个用于读，一个用于写（就是readLoop和writeLoop）。
//在大多数情况下，readLoop会检测socket是否关闭，并适时关闭connection。
//如果一个新请求在readLoop检测到关闭之前就到来了，那么就会产生EOF错误并中断执行，而不是去关闭前一个请求。
//这里也是如此，我执行时建立一个新的连接，这段程序执行完成后退出，再次打开执行时服务器并不知道我已经关闭了连接，
//所以提示连接被重置；如果我不退出程序而使用for循环多次发送时，旧连接未关闭，新连接却到来，会报EOF。