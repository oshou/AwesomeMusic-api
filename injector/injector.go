package injector

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/oshou/AwesomeMusic-api/domain/repository"
	"github.com/oshou/AwesomeMusic-api/infrastructure/datastore/postgres"
	"github.com/oshou/AwesomeMusic-api/presenter/handler"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

type IInjector interface {
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

type injector struct {
	conn *sqlx.DB
}

var _ IInjector = (*injector)(nil)

func NewInjector(conn *sqlx.DB) IInjector {
	return &injector{
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
func (i *injector) NewAppHandler() handler.IAppHandler {
	appHandler := &appHandler{}
	appHandler.IUserHandler = i.NewUserHandler()
	appHandler.ICommentHandler = i.NewCommentHandler()
	appHandler.IPostHandler = i.NewPostHandler()
	appHandler.ITagHandler = i.NewTagHandler()
	appHandler.ISearchHandler = i.NewSearchHandler()

	return appHandler
}

// User
func (i *injector) NewUserRepository() repository.IUserRepository {
	return postgres.NewUserRepository(i.conn)
}

func (i *injector) NewUserUsecase() usecase.IUserUsecase {
	return usecase.NewUserUsecase(i.NewUserRepository())
}

func (i *injector) NewUserHandler() handler.IUserHandler {
	return handler.NewUserHandler(i.NewUserUsecase())
}

// Comment
func (i *injector) NewCommentRepository() repository.ICommentRepository {
	return postgres.NewCommentRepository(i.conn)
}

func (i *injector) NewCommentUsecase() usecase.ICommentUsecase {
	return usecase.NewCommentUsecase(i.NewCommentRepository())
}

func (i *injector) NewCommentHandler() handler.ICommentHandler {
	return handler.NewCommentHandler(i.NewCommentUsecase())
}

// Post
func (i *injector) NewPostRepository() repository.IPostRepository {
	return postgres.NewPostRepository(i.conn)
}

func (i *injector) NewPostUsecase() usecase.IPostUsecase {
	return usecase.NewPostUsecase(i.NewPostRepository())
}

func (i *injector) NewPostHandler() handler.IPostHandler {
	return handler.NewPostHandler(i.NewPostUsecase())
}

// Tag
func (i *injector) NewTagRepository() repository.ITagRepository {
	return postgres.NewTagRepository(i.conn)
}

func (i *injector) NewTagUsecase() usecase.ITagUsecase {
	return usecase.NewTagUsecase(i.NewTagRepository())
}

func (i *injector) NewTagHandler() handler.ITagHandler {
	return handler.NewTagHandler(i.NewTagUsecase())
}

// Search
func (i *injector) NewSearchRepository() repository.ISearchRepository {
	return postgres.NewSearchRepository(i.conn)
}

func (i *injector) NewSearchUsecase() usecase.ISearchUsecase {
	return usecase.NewSearchUsecase(i.NewSearchRepository())
}

func (i *injector) NewSearchHandler() handler.ISearchHandler {
	return handler.NewSearchHandler(i.NewSearchUsecase())
}
