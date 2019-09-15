package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oshou/AwesomeMusic-api/controller"
)

// APIサーバ起動
func Init() {
	r := router()
	r.Run(":" + os.Getenv("API_SERVER_PORT"))
}

func SetCorsPolicy() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			// for preflight
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			ctx.Data(200, "text/plain", []byte{})
			ctx.Abort()
		} else {
			// for actual response
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Next()
		}

		return
	}
}

// ルーティング定義
func router() *gin.Engine {
	r := gin.Default()

	r.Use(SetCorsPolicy())

	u := r.Group("/v1/")
	{
		ctrl := controller.Controller{}

		// ユーザー一覧
		u.GET("/users", ctrl.GetUsers)
		// ユーザー新規追加
		u.POST("/users", ctrl.AddUser)
		// 特定ユーザー表示
		u.GET("/users/:user_id", ctrl.GetUserById)

		// 投稿一覧
		u.GET("/posts", ctrl.GetPosts)
		// 投稿新規追加
		u.POST("/posts", ctrl.AddPost)
		// 特定投稿表示
		u.GET("/posts/:post_id", ctrl.GetPostById)
		// 特定タグの投稿一覧
		u.GET("/tags/:tag_id/posts", ctrl.GetPostsByTagId)
		// 特定ユーザーの投稿一覧
		u.GET("/users/:user_id/posts", ctrl.GetPostsByUserId)
		// // 投稿削除
		// u.DELETE("/posts/:post_id", ctrl.DeletePostById)

		// コメント一覧
		u.GET("/posts/:post_id/comments", ctrl.GetComments)
		// コメント新規追加
		u.POST("/posts/:post_id/comments", ctrl.AddComment)
		// 特定コメント表示
		u.GET("/posts/:post_id/comments/:comment_id", ctrl.GetCommentById)
		// // コメント削除
		// u.DELETE("/posts/:post_id/comments/:comment_id", ctrl.DeleteComment)

		// タグ一覧
		u.GET("/tags", ctrl.GetTags)
		// タグ新規追加
		u.POST("/tags", ctrl.AddTag)
		// タグ表示(特定ID)
		u.GET("/tags/:tag_id", ctrl.GetTagById)
		// 投稿に紐付いているタグ一覧
		u.GET("/posts/:post_id/tags", ctrl.GetTagsByPostId)
		// 投稿へのタグ付与
		u.POST("/posts/:post_id/tags/:tag_id", ctrl.AttachTag)

		// 検索結果
		u.GET("/search", ctrl.SearchByType)
	}

	return r
}
