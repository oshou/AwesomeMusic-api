package handler

type IAppHandler interface {
	IUserHandler
	ICommentHandler
	IPostHandler
	ITagHandler
	ISearchHandler
}
