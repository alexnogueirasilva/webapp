package models

// Publication stores a publication data
type Publication struct {
	ID         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorID   uint64 `json:"author_id,omitempty"`
	AuthorNick string `json:"author_nick,omitempty"`
	Likes      uint64 `json:"likes"`
	CreatedAt  string `json:"created_at,omitempty"`
}
