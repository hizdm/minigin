package models

import (
	"fmt"
	orm "minigin/dao"
)

type Article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// 创建文章内容
func (article *Article) CreateArticle() (id int64, err error) {
	result := orm.Db.Create(&article)

	if result.Error != nil {
		fmt.Println("error err:%v", result.Error)
	}

	return article.ID, result.Error
}

// 查询文章信息
func (article *Article) FindArticle() (articles []Article, err error) {
	if err = orm.Db.Debug().Select("id", "title", "content").Find(&articles).Error; err != nil {
		return
	}

	return
}
