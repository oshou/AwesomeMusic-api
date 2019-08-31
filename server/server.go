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

func SetCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Max-Age", "86400")
		ctx.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}

// ルーティング定義
func router() *gin.Engine {
	r := gin.Default()
	
	r.Use(SetCors())

	u := r.Group("/v1/")
	{
		ctrl := controller.Controller{}

		// ユーザー一覧
		u.GET("/users", ctrl.GetUsers)
		// ユーザー追加
		u.POST("/users", ctrl.AddUser)
		// ユーザー表示
		u.GET("/users/:user_id", ctrl.GetUserById)

		// 投稿一覧
		u.GET("/posts", ctrl.GetPosts)
		// 投稿追加
		u.POST("/posts", ctrl.AddPost)
		// 投稿表示(特定ID)
		u.GET("/posts/:post_id", ctrl.GetPostById)
		// // 投稿削除
		// u.DELETE("/posts/:post_id", ctrl.DeletePostById)

		// コメント一覧
		u.GET("/posts/:post_id/comments", ctrl.GetComments)
		// コメント追加
		u.POST("/posts/:post_id/comments", ctrl.AddComment)
		// コメント表示
		u.GET("/posts/:post_id/comments/:comment_id", ctrl.GetCommentById)
		// // コメント削除
		// u.DELETE("/posts/:post_id/comments/:comment_id", ctrl.DeleteComment)

		// タグ一覧
		u.GET("/tags", ctrl.GetTags)
		// タグ追加
		u.POST("/tags", ctrl.AddTag)
		// タグ表示(特定ID)
		u.GET("/tags/:tag_id", ctrl.GetTagById)
		
		// // 検索結果
		// u.GET("/search", ctrl.SearchByType)
	}

	return r
}
