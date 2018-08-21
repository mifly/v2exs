package v2exs

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
)
type ResultBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type TopicsResult struct {
	ResultBase
	Topics []Topic
}

type Topic struct {
	Title       string
	Url         string
	Category    string
	CategoryUrl string
	MemberName  string
	MemberUrl   string
	MemberImg   string
	LastTime    string
	AnswerCount string
}

func GetTopics(c echo.Context) error {
	fmt.Printf("get topics \n")
	t := c.Param("t")
	topics := GetTopicsByTag(t)
	return c.JSON(http.StatusOK, TopicsResult{ResultBase:ResultBase{Code:0,Msg:"ok"},Topics:topics})
}

func GetTopicsByTag(t string) []Topic {
	var err error
	fmt.Printf("get topic %s begin \n", t)
	resp, err := http.Get("https://www.v2ex.com/?tab=" + t)
	if err != nil {
		fmt.Printf("get latest failed %s\n", err.Error())
		return []Topic{}
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("get hot failed %d\n", resp.StatusCode)
		return []Topic{}
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Printf("get hot failed %d\n", resp.StatusCode)
		return []Topic{}
	}
	var topics []Topic
	topics = make([]Topic, 0)
	doc.Find("#Main .cell.item ").Each(func(i int, s *goquery.Selection) {
		titleEleP := s.Find(".item_title")
		titleEle := titleEleP.Find("a")
		href, _ := titleEle.Attr("href")
		title := titleEle.Text()
		topicInfoEle := s.Find(".topic_info")
		topicInfo := topicInfoEle.Text()
		nodeEle := topicInfoEle.Find(".node")
		cateUrl, _ := nodeEle.Attr("href")
		cate := nodeEle.Text()
		memberName := topicInfoEle.Find("strong a").Text()
		memEle := s.Find("table tr td a")
		membrImg, _ := memEle.Find("img").Attr("src")
		memberHerf, _ := memEle.Attr("href")
		count := s.Find(".count_livid").Text()

		//fmt.Printf("Review %s %s %s \n", href, memberHerf, membrImg)
		t := Topic{Title: title, Url: href, Category: cate, CategoryUrl: cateUrl, MemberName: memberName,
			MemberUrl: memberHerf, MemberImg: membrImg, AnswerCount: count, LastTime: topicInfo}
		topics = append(topics, t)
	})
	//doc.Find(".item_title").Each(func(i int, s *goquery.Selection) {
	//	// For each item found, get the band and title
	//	titleEle := s.Find("a")
	//	href,_ := titleEle.Attr("href")
	//	fmt.Printf("Review %s- %s\n", href,  titleEle.Text())
	//})
	fmt.Println("get topic done")
	////*[@id="Main"]/div[2]/div[2]
	//#Main > div:nth-child(2) > div:nth-child(2)
	return topics
}

