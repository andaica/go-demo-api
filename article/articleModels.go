package article

import (
	"time"
)

type Article struct {
	Id          uint      `json:"id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AuthorId    uint      `json:"authorId"`
}

type Author struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
}

type ArticleResponse struct {
	Article
	Author Author `json:"author"`
}

type Request struct {
	Article Article `json:"article"`
}

type Response struct {
	Article ArticleResponse `json:"article"`
	Status  string          `json:"status"`
}

var Articles = []Article{
	Article{1, "how-to-train-your-dragon", "How to train your dragon", "Ever wonder how?", "It takes a Jacobian", time.Time{}, time.Time{}, 1},
	Article{2, "how-to-train-your-dragon-2", "How to train your dragon 2", "So toothless", "It a dragon", time.Time{}, time.Time{}, 1},
}

var AuthorDemo = Author{"Jake", "jake@test.com", "I work at statefarm", nil}
