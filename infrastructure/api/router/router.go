package router

import "github.com/gin-gonic/gin"

func NewRouter(c *gin.Context) {
	// ユーザー一覧
	u.GET("/users", handler.GetUsers)
	// ユーザー新規追加
	u.POST("/users", ctrl.AddUser)
	// 特定ユーザー表示
	u.GET("/users/:user_id", ctrl.GetUserByID)

	// 投稿一覧
	u.GET("/posts", ctrl.GetPosts)
	// 投稿新規追加
	u.POST("/posts", ctrl.AddPost)
	// 特定投稿表示
	u.GET("/posts/:post_id", ctrl.GetPostByID)
	// 特定タグの投稿一覧
	u.GET("/tags/:tag_id/posts", ctrl.GetPostsByTagID)
	// 特定ユーザーの投稿一覧
	u.GET("/users/:user_id/posts", ctrl.GetPostsByUserID)
	// // 投稿削除
	// u.DELETE("/posts/:post_id", ctrl.DeletePostByID)

	// コメント一覧
	u.GET("/posts/:post_id/comments", ctrl.GetComments)
	// コメント新規追加
	u.POST("/posts/:post_id/comments", ctrl.AddComment)
	// 特定コメント表示
	u.GET("/posts/:post_id/comments/:comment_id", ctrl.GetCommentByID)
	// // コメント削除
	// u.DELETE("/posts/:post_id/comments/:comment_id", ctrl.DeleteComment)

	// タグ一覧
	u.GET("/tags", ctrl.GetTags)
	// タグ新規追加
	u.POST("/tags", ctrl.AddTag)
	// タグ表示(特定ID)
	u.GET("/tags/:tag_id", ctrl.GetTagByID)
	// 投稿に紐付いているタグ一覧
	u.GET("/posts/:post_id/tags", ctrl.GetTagsByPostID)
	// 投稿へのタグ付与
	u.POST("/posts/:post_id/tags/:tag_id", ctrl.AttachTag)

	// 検索結果
	u.GET("/search", ctrl.SearchByType)
}
