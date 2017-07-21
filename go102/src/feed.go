package main

type Feed struct {
	FeedId    int
	Title     string
	Content   string
	Url       string
	Sort      int
	Visible   string
	FeedType  string
	CreatedAt string
	UpdatedAt *string
}

func Feeds() (feeds []Feed, err error) {

	rows, err := DB.Query("SELECT feed_id, title, content, url, sort, visible, feed_type, created_at, updated_at FROM feeds")

	if err == nil {
		defer rows.Close()

		for rows.Next() {
			feed := Feed{}
			rows.Scan(&feed.FeedId, &feed.Title, &feed.Content, &feed.Url, &feed.Sort, &feed.Visible, &feed.FeedType, &feed.CreatedAt, &feed.UpdatedAt)
			feeds = append(feeds, feed)
		}

	}
	return

}

func SelectFeedById(id int) (feed Feed, err error) {
	feed = Feed{}
	err = DB.QueryRow("SELECT feed_id, title, content, url, sort, visible, feed_type, created_at, updated_at FROM feeds WHERE feed_id = ?", id).Scan(&feed.FeedId, &feed.Title, &feed.Content, &feed.Url, &feed.Sort, &feed.Visible, &feed.FeedType, &feed.CreatedAt, &feed.UpdatedAt)
	return
}
