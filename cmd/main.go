package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"novaro-server/api"
	"novaro-server/config"
	"novaro-server/model"
)

func init() {
	config.Init()
	db := config.DB
	db.AutoMigrate(&model.Collections{})
	db.AutoMigrate(&model.Comments{})
	db.AutoMigrate(&model.Posts{})
	db.AutoMigrate(&model.RePosts{})
	db.AutoMigrate(&model.Tags{})
	db.AutoMigrate(&model.TagsRecords{})
	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.TwitterUser{})
}
func main() {

	// 创建 cron 实例
	c := cron.New()

	// 添加定时任务：每小时刷新 Redis 键的过期时间
	c.AddFunc("@every 1m", func() {
		model.SyncData()
	})

	c.Start()

	r := gin.Default()
	// 使用gin-swagger中间件来提供API文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	collections := r.Group("/api/collections")
	{
		collections.POST("/add", api.CollectionsApi{}.CollectionsTweet)
		collections.POST("/remove", api.CollectionsApi{}.UnCollectionsTweet)
	}

	comments := r.Group("/api/comments")
	{
		comments.GET("/getCommentsListByPostId", api.CommentsApi{}.GetCommentsListByPostId)
		comments.GET("/getCommentsListByParentId", api.CommentsApi{}.GetCommentsListByParentId)
		comments.GET("/getCommentsListByUserId", api.CommentsApi{}.GetCommentsListByUserId)
		comments.POST("/add", api.CommentsApi{}.AddComments)
	}

	posts := r.Group("/api/posts")
	{
		posts.GET("/getPostsById", api.PostsApi{}.GetPostsById)
		posts.GET("/getPostsByUserId", api.PostsApi{}.GetPostsByUserId)
		posts.GET("/getPostsList", api.PostsApi{}.GetPostsList)
		posts.POST("/savePosts", api.PostsApi{}.SavePosts)
		posts.DELETE("/delPostsById", api.PostsApi{}.DelPostsById)
	}

	reposts := r.Group("/api/reposts")
	{
		reposts.POST("/add", api.RePostsApi{}.AddRePosts)
	}

	tags := r.Group("/api/tags")
	{
		tags.GET("/getTagsList", api.TagsApi{}.GetTagsList)
	}

	records := r.Group("/api/tags/records")
	{
		records.GET("/add", api.TagsRecordsApi{}.AddTagsRecords)
	}

	r.Run()
}
