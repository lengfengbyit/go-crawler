package engine

// 存放请求URL， 和内容解析函数
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

// 存放解析结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

