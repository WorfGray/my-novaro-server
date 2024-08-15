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
	r.Group("/api/collections")
	{
		r.POST("/add", api.CollectionsApi{}.CollectionsTweet)
		r.POST("/remove", api.CollectionsApi{}.UnCollectionsTweet)
	}

	r.Group("/api/comments")
	{
		r.GET("/getCommentsListByPostId", api.CommentsApi{}.GetCommentsListByPostId)
		r.GET("/getCommentsListByParentId", api.CommentsApi{}.GetCommentsListByParentId)
		r.GET("/getCommentsListByUserId", api.CommentsApi{}.GetCommentsListByUserId)
		r.POST("/add", api.CommentsApi{}.AddComments)
	}

	r.Group("/api/posts")
	{
		r.GET("/getPostsById", api.PostsApi{}.GetPostsById)
		r.GET("/getPostsByUserId", api.PostsApi{}.GetPostsByUserId)
		r.GET("/getPostsList", api.PostsApi{}.GetPostsList)
		r.POST("/savePosts", api.PostsApi{}.SavePosts)
		r.DELETE("/delPostsById", api.PostsApi{}.DelPostsById)
	}

	r.Group("/api/reposts")
	{
		r.POST("/add", api.RePostsApi{}.AddRePosts)
	}

	r.Group("/api/tags")
	{
		r.GET("/getTagsList", api.TagsApi{}.GetTagsList)
	}

	r.Group("/api/tags/records")
	{
		r.GET("/add", api.TagsRecordsApi{}.AddTagsRecords)
	}

	r.Run()
}
