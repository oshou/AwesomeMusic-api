// Package handler is ui layer http-handler package
package handler

// IAppHandler is aggregation handler interface
type IAppHandler interface {
	IUserHandler
	ICommentHandler
	IPostHandler
	ITagHandler
	ISearchHandler
}
