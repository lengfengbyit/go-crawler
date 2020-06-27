package parse

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"log"
	"project/crawl/engine"
	"project/crawl/model"
)

func GetBookDetail(content []byte) engine.ParseResult {

	//books := []model.Book{}
	html, _ := htmlquery.Parse(bytes.NewReader(content))
	nameNode := htmlquery.FindOne(html, "//*[@id=\"wrapper\"]/h1/span")
	autherNode := htmlquery.FindOne(html, "//*[@id=\"info\"]/span[1]/a")

	book := model.Book{
		Name:   htmlquery.InnerText(nameNode),
		Author: htmlquery.InnerText(autherNode),
	}

	log.Printf("书名: %s, 作者: %s\n", book.Name, book.Author)
	return engine.ParseResult{}
}
