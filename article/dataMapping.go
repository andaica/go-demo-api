package article

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/andaica/go-demo-api/db"

	_ "github.com/go-sql-driver/mysql"
)

const table = "article"

type DataMapping struct{}

func (a DataMapping) fetchAll() []ArticleResponse {
	articles := []ArticleResponse{}
	query := `select article.*, user.username, user.email, user.bio, user.image
		from article left join user on article.authorId = user.id`
	result, _ := db.Execute(query)

	for result.Next() {
		article := scanToArticleResponse(result)
		articles = append(articles, article)
	}

	return articles
}

func (a DataMapping) insertNewArticle(article Article, userId uint) (newArt ArticleResponse, isOk bool) {
	query := fmt.Sprintf(
		"insert into article ( slug, title, description, body, authorId ) values ( '%s', '%s', '%s', '%s', %d )",
		article.Slug, article.Title, article.Description, article.Body, userId,
	)
	res, err := db.ExecuteOne(query)

	if err == nil {
		id, err := res.LastInsertId()
		if err == nil {
			newArt, isExist := getArticleById(uint(id))
			return newArt, isExist
		}
	}

	return newArt, false
}

func getArticleById(id uint) (newArticle ArticleResponse, isExist bool) {
	query := `select article.*, user.username, user.email, user.bio, user.image
		from article left join user on article.authorId = user.id
		where article.id = ` + strconv.Itoa(int(id))
	result, _ := db.Execute(query)

	if result.Next() {
		newArticle = scanToArticleResponse(result)
		return newArticle, true
	}

	return newArticle, false
}

func scanToArticleResponse(row *sql.Rows) ArticleResponse {
	art := Article{}
	author := Author{}
	err := row.Scan(
		&art.Id, &art.Slug, &art.Title, &art.Description, &art.Body, &art.CreatedAt, &art.UpdatedAt, &art.AuthorId,
		&author.Username, &author.Email, &author.Bio, &author.Image,
	)
	if err != nil {
		panic(err.Error())
	}

	result := ArticleResponse{
		Article: art,
		Author:  author,
	}
	return result
}
