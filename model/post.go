package model

import (
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Post is a corresponding to a post on ptt
type Post struct {
	Title string
	Href  string
	NVote int
	Date  time.Time
}

// fetchPreviewImg get the preview image of a post
func fetchPreviewImg(p *Post) string {
	// TODO: handle error
	doc, _ := goquery.NewDocument(p.Href)
	imgSelector := `#main-content a[href$=".jpg"],a[href$=".png"],a[href$=".gif"],a[href*="imgur"]`
	imgURL, _ := doc.Find(imgSelector).Attr("href")

	// https://imgur.com/8bsl72C -> https://imgur.com/8bsl72C.jpg
	matched, err := regexp.MatchString("^https?://imgur.com/\\w+$", imgURL)
	if err != nil {
		panic(err)
	}
	if matched {
		imgURL += ".jpg"
	}

	return imgURL
}

// [正妹] 大橋未久 -> 大橋未久
func trimTitlePrefix(title string) string {
	return strings.TrimPrefix(title, "[正妹] ")
}

// ToBeauty transform a Post to a Beauty
func (p *Post) ToBeauty() Beauty {
	previewImg := fetchPreviewImg(p)
	return Beauty{
		NVote:      p.NVote,
		Title:      trimTitlePrefix(p.Title),
		Href:       p.Href,
		PreviewImg: previewImg,
	}
}