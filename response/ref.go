package response

// ResponseItem は、Responseのリストアイテムの構造体。
type ResponseItem struct {
	ID            string         `json:"id"`
	ArticleID     string         `json:"article_id"`
	ParentID      string         `json:"parent_id"`
	UserName      string         `json:"user_name"`
	Content       string         `json:"content"`
	CreatedAt     string         `json:"created_at"`
	ReplyComments []ResponseItem `json:"reply_comments"`
}
