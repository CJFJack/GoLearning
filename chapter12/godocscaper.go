package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	q := "goquery"
	url := "https://godoc.org/?q=" + q

	// 发送http请求，并创建document结构体指针
	document, err := goquery.NewDocument(url)
	fmt.Println(document, err)

	// class选择器
	selection := document.Find(".table-condensed")

	// tag选择器
	selection = selection.Find("a")
	selection.Each(func(i int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")
		if !exists {
			fmt.Println("href attr is not exists")
		}
		fmt.Println(selection.Text(), href, exists)
	})

	// id选择器
	fmt.Println(document.Find("#x-search").Attr("class"))
	fmt.Println(document.Find("#x-search").Html())
	fmt.Println(document.Find("#x-search").Text())

	// 复合选择器
	selection = document.Find("table.table-condensed")
	fmt.Println(selection.Text())

	// 子孙选择器
	document.Find(".table-condensed a").Each(func(i int, selection *goquery.Selection) {
		href, exists := selection.Attr("href")
		if !exists {
			fmt.Println("href attr is not exists")
		}
		fmt.Println(selection.Text(), href, exists)
	})
}
