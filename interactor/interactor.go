package interactor

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/infrastructure/datastore"
	"github.com/oshou/AwesomeMusic-api/presenter/handler"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

type IInteractor interface {
	// App
	NewAppHandler() handler.IAppHandler
	// User
	NewUserRepository() repository.IUserRepository
	NewUserUsecase() usecase.IUserUsecase
	NewUserHandler() handler.IUserHandler
	// Comment
	NewCommentRepository() repository.ICommentRepository
	NewCommentUsecase() usecase.ICommentUsecase
	NewCommentHandler() handler.ICommentHandler
	// Post
	NewPostRepository() repository.IPostRepository
	NewPostUsecase() usecase.IPostUsecase
	NewPostHandler() handler.IPostHandler
	// Tag
	NewTagRepository() repository.ITagRepository
	NewTagUsecase() usecase.ITagUsecase
	NewTagHandler() handler.ITagHandler
	// Search
	NewSearchRepository() repository.ISearchRepository
	NewSearchUsecase() usecase.ISearchUsecase
	NewSearchHandler() handler.ISearchHandler
}

type interactor struct {
	conn *sqlx.DB
}

var _ IInteractor = (*interactor)(nil)

func NewInteractor(conn *sqlx.DB) IInteractor {
	return &interactor{
		conn: conn,
	}
}

type appHandler struct {
	handler.IUserHandler
	handler.ICommentHandler
	handler.IPostHandler
	handler.ITagHandler
	handler.ISearchHandler
}

// Aggregate Handler
func (i *interactor) NewAppHandler() handler.IAppHandler {
	appHandler := &appHandler{}
	appHandler.IUserHandler = i.NewUserHandler()
	appHandler.ICommentHandler = i.NewCommentHandler()
	appHandler.IPostHandler = i.NewPostHandler()
	appHandler.ITagHandler = i.NewTagHandler()
	appHandler.ISearchHandler = i.NewSearchHandler()

	return appHandler
}

// User
func (i *interactor) NewUserRepository() repository.IUserRepository {
	return datastore.NewUserRepository(i.conn)
}

func (i *interactor) NewUserUsecase() usecase.IUserUsecase {
	return usecase.NewUserUsecase(i.NewUserRepository())
}

func (i *interactor) NewUserHandler() handler.IUserHandler {
	return handler.NewUserHandler(i.NewUserUsecase())
}

// Comment
func (i *interactor) NewCommentRepository() repository.ICommentRepository {
	return datastore.NewCommentRepository(i.conn)
}

func (i *interactor) NewCommentUsecase() usecase.ICommentUsecase {
	return usecase.NewCommentUsecase(i.NewCommentRepository())
}

func (i *interactor) NewCommentHandler() handler.ICommentHandler {
	return handler.NewCommentHandler(i.NewCommentUsecase())
}

// Post
func (i *interactor) NewPostRepository() repository.IPostRepository {
	return datastore.NewPostRepository(i.conn)
}

func (i *interactor) NewPostUsecase() usecase.IPostUsecase {
	return usecase.NewPostUsecase(i.NewPostRepository())
}

func (i *interactor) NewPostHandler() handler.IPostHandler {
	return handler.NewPostHandler(i.NewPostUsecase())
}

// Tag
func (i *interactor) NewTagRepository() repository.ITagRepository {
	return datastore.NewTagRepository(i.conn)
}

func (i *interactor) NewTagUsecase() usecase.ITagUsecase {
	return usecase.NewTagUsecase(i.NewTagRepository())
}

func (i *interactor) NewTagHandler() handler.ITagHandler {
	return handler.NewTagHandler(i.NewTagUsecase())
}

// Search
func (i *interactor) NewSearchRepository() repository.ISearchRepository {
	return datastore.NewSearchRepository(i.conn)
}

func (i *interactor) NewSearchUsecase() usecase.ISearchUsecase {
	return usecase.NewSearchUsecase(i.NewSearchRepository())
}

func (i *interactor) NewSearchHandler() handler.ISearchHandler {
	return handler.NewSearchHandler(i.NewSearchUsecase())
}
