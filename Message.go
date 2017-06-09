package main

// Message ...
type Message struct {
	ID          string   `json:"id"`
	Text        string   `json:"text"`
	AuthorID    string   `json:"authorId"`
	RecipientID string   `json:"recipientId"`
	Likes       []string `json:"likes"`
	Created     string   `json:"created"`
	Edited      string   `json:"edited"`
}
