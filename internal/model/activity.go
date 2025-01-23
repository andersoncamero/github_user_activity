package model

type Activity struct {
	Id        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	Public    bool    `json:"public"`
	CreatedAt string  `json:"created_at"`
}

type Actor struct {
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
}

type Repo struct {
	Name string `json:"name"`
}

type Payload struct {
	PullRequest PullRequest `json:"pull_request"`
	Description string      `json:"description"`
}

type PullRequest struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}
