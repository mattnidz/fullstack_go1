package actions

import (
	"fmt"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gocolly/colly"
)

type Story struct {
	Rank  string
	Title string
	URL   string
}

type topStoriesSnapshot struct {
	Stories []Story
	FoundAt time.Time
}

func CreateStory(e *colly.HTMLElement) Story {
	rank := e.ChildText(".rank")
	url := e.ChildAttr(".storylink", "href")
	title := e.ChildText(".storylink")

	return Story{Rank: rank, URL: url, Title: title}
}

// NewsList default implementation.
func NewsList(c buffalo.Context) error {
	a := colly.NewCollector()
	snapshot := topStoriesSnapshot{FoundAt: time.Now()}

	a.OnHTML(".athing", func(e *colly.HTMLElement) {
		story := CreateStory(e)
		snapshot.Stories = append(snapshot.Stories, story)
	})

	a.Visit("https://news.ycombinator.com")
	fmt.Println(snapshot)

	//return c.Render(200, r.HTML("news/list.html"))
	return c.Render(200, r.JSON(snapshot))

}
