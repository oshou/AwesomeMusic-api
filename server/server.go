package server

import (
	"github.com/oshou/AwesomeMusic-api/controller"
	"github.com/gin-gonic/gin"
)

// APIサーバ起動
func Init() {
	r := router()
	r.Run()
}

// ルーティング定義
func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/v1/")
	{
		ctrl := task.Controller{}

		// ユーザー一覧
		u.GET("/users", ctrl.GetUsers)
		// ユーザー追加
		u.POST("/users", ctrl.AddUser)
		// 投稿一覧
		u.GET("/posts", ctrl.GetPosts)
		// 投稿追加
		u.POST("/posts", ctrl.AddPost)
		// 投稿表示
		u.GET("/posts/:post_id", ctrl.GetPost)
		// 投稿削除
		u.DELETE("/posts/:post_id", ctrl.DeletePost)
		// コメント一覧
		u.GET("/posts/:post_id/comments", ctrl.GetComments)
		// コメント追加
		u.POST("/posts/:post_id/comments", ctrl.AddComment)
		// コメント表示
		u.GET("/posts/:post_id/comments/:comment_id", ctrl.GetComment)
		// コメント削除
		u.DELETE("/posts/:post_id/comments/:comment_id", ctrl.DeleteComment)
		// タグ一覧
		u.GET("/tags", ctrl.GetTags)
		// 検索結果
		u.GET("/search", ctrl.SearchByType)
	}

	return r
}
