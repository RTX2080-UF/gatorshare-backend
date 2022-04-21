package models

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
	  q := r.URL.Query()
	  page, _ := strconv.Atoi(q.Get("page"))
	  if page == 0 {
		page = 1
	  }
  
	  pageSize, _ := strconv.Atoi(q.Get("page_size"))
	  switch {
	  case pageSize > 100:
		pageSize = 100
	  case pageSize <= 0:
		pageSize = 10
	  }
  
	  offset := (page - 1) * pageSize
	  return db.Offset(offset).Limit(pageSize)
	}
}

func GetUserHomePosts(db *gorm.DB, id uint) ([]TagPost, error) {
	var userLikedTags []TagUser
	res := db.Where("user_id = ?", id).Find(&userLikedTags);

	var tagIDs []uint
	for  _, entity := range userLikedTags {
		tagIDs = append(tagIDs, entity.TagID)
	}
	
	var tagPosts []TagPost
	res = db.Preload("Post").Preload("Post.User").Distinct("post_id").Order("created_at desc, updated_at desc").Where("tag_id IN(?)", tagIDs).Find(&tagPosts)
	return tagPosts, res.Error
}

func GetLatestPost(db *gorm.DB, ctx *gin.Context, post *[]Post, id uint) (error) {
	res := db.Preload("User").Scopes(Paginate(ctx.Request)).Order("created_at desc, updated_at desc").Find(&post)
	return res.Error
}