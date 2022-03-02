package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mtanzim/gopl-book/ch4/github"
)

const (
	LTMonth = iota
	LTYear
	MTYear
)

var Messages = map[int]string{
	LTMonth: "Less than a month",
	LTYear:  "Less than a year",
	MTYear:  "More than a year",
}

func categorize(item *github.Issue) int {
	timeSinceCreated := time.Since(item.CreatedAt)

	if timeSinceCreated < time.Hour*24*30 {
		return LTMonth
	}
	if timeSinceCreated < time.Hour*24*365 {
		return LTYear
	}
	return MTYear

}

func makeCategoryMap(items []*github.Issue) map[int][]*github.Issue {
	itemCategoryMap := make(map[int][]*github.Issue)

	for _, item := range items {
		category := categorize(item)

		current := itemCategoryMap[category]
		itemCategoryMap[category] = append(current, item)
	}

	return itemCategoryMap

}
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	itemCategoryMap := makeCategoryMap(result.Items)
	for category, values := range itemCategoryMap {
		fmt.Printf("\n%s\n\n", Messages[category])
		for _, item := range values {
			category := categorize(item)
			fmt.Printf("#%-5d %s\t\t %9.9s %.55s\n",
				item.Number, Messages[category], item.User.Login, item.Title)
		}
	}

}
