package book

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	book2 "project/crawl/model/douban/book"
	"project/crawl/scheduler"
	"project/crawl/util"
	"regexp"
	"strconv"
	"strings"
)

func GetBookDetail(content []byte) scheduler.ParseResult {

	//books := []model.Book{}
	html, _ := htmlquery.Parse(bytes.NewReader(content))
	nameNode := htmlquery.FindOne(html, "//*[@id=\"wrapper\"]/h1/span")
	authorNode := htmlquery.FindOne(html, "//*[@id=\"info\"]//a")
	coverNode := htmlquery.FindOne(html, "//*[@id=\"mainpic\"]/a/img")
	pressNode := htmlquery.FindOne(html, "//*[@id=\"info\"]]")

	starNode := htmlquery.FindOne(html, "//*[@id=\"interest_sectl\"]/div/div[2]/div/div[1]")
	starClass := htmlquery.SelectAttr(starNode, "class")
	star, _ := strconv.Atoi(starClass[strings.Index(starClass, "bigstar")+7:])

	evaluationNumNode := htmlquery.FindOne(html, "//*[@class=\"rating_people\"]/span")
	evaluationNum := 0
	if evaluationNumNode != nil {
		evaluationNum, _ = strconv.Atoi(htmlquery.InnerText(evaluationNumNode))
	}

	scoreNode := htmlquery.FindOne(html, "//*[@id=\"interest_sectl\"]/div/div[2]/strong")
	score, _ := strconv.ParseFloat(strings.Trim(htmlquery.InnerText(scoreNode), " "), 1)

	introNode := htmlquery.FindOne(html, "//*[@class=\"intro\"]")
	intro := ""
	if introNode != nil {
		intro = strings.TrimSpace(htmlquery.InnerText(introNode))
	}

	author := ""
	if authorNode != nil {
		author = util.TrimSpaceAndLinefeed(htmlquery.InnerText(authorNode))
	}

	metaNode := htmlquery.FindOne(html, "/html/head/meta[4]")
	metaContent := htmlquery.SelectAttr(metaNode, "content")
	ls := strings.Split(metaContent, "/")
	id, err := strconv.Atoi(ls[len(ls)-2])
	if err != nil {
		id = 0
	}

	book := book2.Book{
		Id:            uint32(id),
		Name:          htmlquery.InnerText(nameNode),
		Author:        author,
		CoverUrl:      htmlquery.SelectAttr(coverNode, "src"),
		Intro:         intro,
		Star:          float32(star) / 10,
		EvaluationNum: uint32(evaluationNum),
		Score:         float32(score),
	}

	text := htmlquery.InnerText(pressNode)
	text = strings.Replace(text, " ", "", -1)
	text = regexp.MustCompile("\\s+").ReplaceAllString(text, "\n")
	lines := strings.Split(text, "\n")

	for _, item := range lines {
		tmp := strings.Split(item, ":")

		switch tmp[0] {
		case "出版社":
			book.Press = tmp[1]
		case "出版年":
			book.PressDate = tmp[1]
		case "页数":
			pageSize, _ := strconv.Atoi(tmp[1])
			book.PageSize = uint16(pageSize)
		case "定价":
			book.Price = tmp[1]
		}

	}

	return scheduler.ParseResult{
		Items:    []interface{}{book},
		Requests: nil,
	}
}
