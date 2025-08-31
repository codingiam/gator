package feed

import (
	"codingiam/gator/internal/database"
	"codingiam/gator/internal/state"
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := 0; i < len(feed.Channel.Item); i++ {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	fmt.Println(feed)

	return &feed, nil
}

func ScrapeFeeds(st *state.State) (*RSSFeed, error) {
	ctx := context.Background()

	ffeed, err := st.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return nil, err
	}

	rss, err := FetchFeed(ctx, ffeed.Url)
	if err != nil {
		return nil, err
	}

	err = st.Db.MarkFeedFetched(ctx,
		database.MarkFeedFetchedParams{ID: ffeed.ID, LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true}})
	if err != nil {
		return rss, err
	}

	return rss, nil
}
