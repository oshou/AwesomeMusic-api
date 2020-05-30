// Package router is http-router package
package router

import (
	"github.com/gin-gonic/gin"

	"github.com/oshou/AwesomeMusic-api/ui/http/handler"
	"github.com/oshou/AwesomeMusic-api/ui/http/middleware"
)

// IRouter is http-router interface
type IRouter interface {
	Run(addr string) error
}

type router struct {
	engine *gin.Engine
}

var _ IRouter = &router{}

// NewRouter is constructor for router
func NewRouter(e *gin.Engine, h handler.IAppHandler) IRouter {
	e.Use(middleware.SetCors(e))
	g := e.Group("/v1/")
	{
		// ユーザー一覧
		g.GET("/users", h.GetUsers)
		// ユーザー新規追加
		g.POST("/users", h.AddUser)
		// 特定ユーザー表示
		g.GET("/users/:user_id", h.GetUserByID)

		// 投稿一覧
		g.GET("/posts", h.GetPosts)
		// 投稿新規追加
		g.POST("/posts", h.AddPost)
		// 特定投稿表示
		g.GET("/posts/:post_id", h.GetPostByID)
		// 特定タグの投稿一覧
		g.GET("/tags/:tag_id/posts", h.GetPostsByTagID)
		// 特定ユーザーの投稿一覧
		g.GET("/users/:user_id/posts", h.GetPostsByUserID)
		// 投稿削除
		g.DELETE("/posts/:post_id", h.DeletePostByID)

		// コメント一覧
		g.GET("/posts/:post_id/comments", h.GetComments)
		// コメント新規追加
		g.POST("/posts/:post_id/comments", h.AddComment)
		// 特定コメント表示
		g.GET("/posts/:post_id/comments/:comment_id", h.GetCommentByID)
		// // コメント削除
		//  // u.DELETE("/posts/:post_id/comments/:comment_id", ctrl.DeleteComment)

		// タグ一覧
		g.GET("/tags", h.GetTags)
		// タグ新規追加
		g.POST("/tags", h.AddTag)
		// タグ表示(特定ID)
		g.GET("/tags/:tag_id", h.GetTagByID)
		// 投稿に紐付いているタグ一覧
		g.GET("/posts/:post_id/tags", h.GetTagsByPostID)
		// 投稿へのタグ付与
		g.POST("/posts/:post_id/tags/:tag_id", h.AttachTag)

		// 検索結果
		g.GET("/search", h.SearchByType)
	}

	return &router{e}
}

func (r *router) Run(addr string) error {
	err := r.engine.Run(addr)
	return err
}
