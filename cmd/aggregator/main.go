package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"time"

	"github.com/mmcdole/gofeed"

	"github.com/codecure/news-aggregator/database"
)

type feedRule struct {
	date    string
	content string
}

func parseRule(rule string) feedRule {
	parsedRule, err := url.ParseQuery(rule)
	if err != nil {
		log.Fatal(err)
	}

	var resultRule = feedRule{}

	if date, ok := parsedRule["date"]; ok {
		resultRule.date = date[0]
	}
	if content, ok := parsedRule["content"]; ok {
		resultRule.content = content[0]
	}
	return resultRule
}

func applyRule(feed gofeed.Feed, rule feedRule) []database.NewsItem {
	result := make([]database.NewsItem, len(feed.Items))
	for _, feedItem := range feed.Items {
		item := database.NewsItem{
			Title:  feedItem.Title,
			Link:   feedItem.Link,
			Author: feedItem.Author.Name,
			GUID:   feedItem.GUID}
		item.Date = getField(feedItem, rule.date)
		item.Content = getField(feedItem, rule.content)
		result = append(result, item)
	}
	return result
}

func getField(v interface{}, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	fieldValue := f.Interface()

	switch v := fieldValue.(type) {
	case time.Time:
		return v.Format(time.RFC3339)
	default:
		return fmt.Sprintf("%s", v)
	}
}

func main() {
	// URL: https://echo.tochka.com/feeds/topics/ru/ Rule: date=UpdatedParsed&content=Content
	// URL: https://habr.com/ru/rss/all/all/?fl=ru Rule: date=PublishedParsed&content=Description

	feedPtr := flag.String("feed", "", "Feed URL")
	rulePtr := flag.String("rule", "", "Parsing rule")
	dbPathPtr := flag.String("db", "", "Directory to store DB file")
	flag.Parse()
	rule := parseRule(*rulePtr)

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(*feedPtr)
	processedFeed := applyRule(*feed, rule)

	db := database.GetDB(*dbPathPtr)
	defer db.Close()

	db.AutoMigrate(&database.NewsItem{})

	for _, item := range processedFeed {
		db.Create(&item)
	}
}
