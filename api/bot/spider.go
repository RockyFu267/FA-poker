package bot

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// FetchJobAnnouncements 爬取招聘公告
func (q *Qianxun) FetchJobAnnouncements(c *gin.Context) {
	fmt.Println("Starting FetchJobAnnouncements")
	resp, err := http.Get("https://rsj.sh.gov.cn/tzpgg_17408/index.html")
	if err != nil {
		fmt.Println("Error fetching the page:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching the page"})
		return
	}
	fmt.Println("Fetching the page was successful")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error closing the body:", err)
		}
	}(resp.Body)

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing the page:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing the page"})
		return
	}
	fmt.Println("Parsing the page was successful")

	var announcements []map[string]string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" && strings.Contains(a.Val, "tzpgg_17408") {
					title := getTitle(n)
					date := getDate(n)
					fullURL := "https://rsj.sh.gov.cn" + a.Val
					announcements = append(announcements, map[string]string{
						"title": title,
						"url":   fullURL,
						"date":  date,
					})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Println("Finished processing the page")
	c.JSON(http.StatusOK, announcements)
}

func getTitle(n *html.Node) string {
	if n.FirstChild != nil {
		return strings.TrimSpace(n.FirstChild.Data)
	}
	return ""
}

func getDate(n *html.Node) string {
	for c := n.NextSibling; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == "span" && strings.Contains(c.Attr[0].Val, "time") {
			return strings.TrimSpace(c.FirstChild.Data)
		}
	}
	return ""
}

// 访问https://www.birdreport.cn/home/record/page.html?param=eyJyZXBvcnRJZCI6IjA1YzAxYjVlLTM4OTQtNDUwZS04ZDQ5LTA1YmZhZjRjZGNlNiIsInRzIjoxNzM2MjM3MDA1MjM1fQ==&sign=d41d8cd98f00b204e9800998ecf8427e 解析页面

func (q *Qianxun) FetchBirdReport(c *gin.Context) {
	url := "https://www.birdreport.cn/home/record/page.html?param=eyJyZXBvcnRJZCI6IjA1YzAxYjVlLTM4OTQtNDUwZS04ZDQ5LTA1YmZhZjRjZGNlNiIsInRzIjoxNzM2MjQzODE1NjIwfQ==&sign=d41d8cd98f00b204e9800998ecf8427e"

	// Create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var htmlContent string

	// Run task
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &htmlContent),
	)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	var result []map[string]string
	var parseTable func(*html.Node)
	parseTable = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "tr" {
			rowData := make(map[string]string)
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "td" {
					key := getKey(c)
					value := getTextContent(c)
					rowData[key] = value
				}
			}
			if len(rowData) > 0 {
				result = append(result, rowData)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parseTable(c)
		}
	}

	parseTable(doc)
	c.JSON(http.StatusOK, result)
}

func getKey(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "data-field" {
			return a.Val
		}
	}
	return ""
}

func getTextContent(n *html.Node) string {
	var sb strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			sb.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return strings.TrimSpace(sb.String())
}
