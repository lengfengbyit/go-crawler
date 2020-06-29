package model

import (
	"fmt"
)

// 书的信息
type Book struct {
	Id        uint32
	CoverUrl  string // 封面图片链接
	Name      string
	Author    string
	Press     string // 出版社
	PressDate string // 出版日期
	PageSize  uint16 // 页数
	Price     string // 价格

	Star          float32 // 星数
	EvaluationNum uint32  // 评价人数
	Score         float32 // 评分
	Intro         string  // 简介
}

// 重载String函数，自定义输出字符串
func (book Book) String() string {
	str := fmt.Sprintf("\nID:%d\n书名: %s\n作者: %s\n出版社:%s\n出版日期:%s\n页数:%v\n价格:%v\n星数:%v\n评价人数:%v\n评分:%v\n简介:%s\n",
		book.Id, book.Name, book.Author, book.Press, book.PressDate,
		book.PageSize, book.Price, book.Star, book.EvaluationNum,
		book.Score, book.Intro)

	return str
}
