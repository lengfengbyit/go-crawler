package model

// 书的信息
type Book struct {
	CoverUrl  string // 封面图片链接
	Name      string
	Author    string
	Press     string  // 出版社
	PressDate string  // 出版日期
	PageSize  uint16  // 页数
	Price     float32 // 价格

	Star          float32 // 星数
	EvaluationNum float32 // 评价人数
	Score         float32 // 评分
	Intro         string  // 简介
}
