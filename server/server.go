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
		// ユーザー表示
		u.GET("/posts/:post_id", ctrl.GetPosts)
		// ユーザー追加
		u.PUT("/posts/:post_id", ctrl.AddPosts)
		// コメント一覧
		u.GET("/posts/:post_id/comments", ctrl.GetComments)
		// コメント追加
		u.PUT("/posts/:post_id/comments", ctrl.AddComment)
		// コメント表示
		u.GET("/posts/:post_id/comments/:comment_id", ctrl.getComment)
		// コメント削除
		u.DELETE("/posts/:post_id/comments/:comment_id", ctrl.getComment)
		// タグ一覧
		u.GET("/tags", ctrl.ShowScraping)
		// 検索結果
		u.GET("/search", ctrl.ShowMail)
	}

	return r
}
