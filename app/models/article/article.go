package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

/**
 * 定义一个结构 Article
 */
type Article struct {
	models.BaseModel
	Title string
	Body  string
}

/**
 * 给 Article 结构添加一个link方法
 */
func (article Article) Link(a Article) string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}
