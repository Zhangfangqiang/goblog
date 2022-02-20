package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"os/user"
	"strconv"
)

/**
 * 定义一个结构 Article
 */
type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`

	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

/**
 * 给 Article 结构添加一个link方法
 */
func (article Article) Link(a Article) string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}

func (article Article) CreatedAtDate() string {
	return article.CreatedAt.Format("2006-01-02")
}
