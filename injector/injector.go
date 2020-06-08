// Package injector is DI Container
package injector

import (
	"github.com/jmoiron/sqlx"
	"github.com/oshou/AwesomeMusic-api/domain/repository"

	persistence "github.com/oshou/AwesomeMusic-api/infrastructure/persistence/postgres"
	//persistence "github.com/oshou/AwesomeMusic-api/infrastructure/persistence/mysql"

	"github.com/oshou/AwesomeMusic-api/service"
	"github.com/oshou/AwesomeMusic-api/ui/http/handler"
	"github.com/oshou/AwesomeMusic-api/ui/http/router"
)

// IInjector is DI Container interface
type IInjector interface {
	// Router
	NewRouter() router.IRouter
	// App
	NewAppHandler() handler.IAppHandler
	// User
	NewUserRepository() repository.IUserRepository
	NewUserService() service.IUserService
	NewUserHandler() handler.IUserHandler
	// Comment
	NewCommentRepository() repository.ICommentRepository
	NewCommentService() service.ICommentService
	NewCommentHandler() handler.ICommentHandler
	// Post
	NewPostRepository() repository.IPostRepository
	NewPostService() service.IPostService
	NewPostHandler() handler.IPostHandler
	// Tag
	NewTagRepository() repository.ITagRepository
	NewTagService() service.ITagService
	NewTagHandler() handler.ITagHandler
	// Search
	NewSearchRepository() repository.ISearchRepository
	NewSearchService() service.ISearchService
	NewSearchHandler() handler.ISearchHandler
}

// injector
type injector struct {
	conn *sqlx.DB
}

var _ IInjector = &injector{}

// NewInjector is constructor for injector
func NewInjector(conn *sqlx.DB) IInjector {
	return &injector{
		conn: conn,
	}
}

// Router
func (i *injector) NewRouter(h i.NewAppHandler) router.IRouter {
	return router.NewRouter(h)
}

// App Aggregate Handler
type appHandler struct {
	handler.IUserHandler
	handler.ICommentHandler
	handler.IPostHandler
	handler.ITagHandler
	handler.ISearchHandler
}

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
	return persistence.NewUserRepository(i.conn)
}

func (i *injector) NewUserService() service.IUserService {
	return service.NewUserService(i.NewUserRepository())
}

func (i *injector) NewUserHandler() handler.IUserHandler {
	return handler.NewUserHandler(i.NewUserService())
}

// Comment
func (i *injector) NewCommentRepository() repository.ICommentRepository {
	return persistence.NewCommentRepository(i.conn)
}

func (i *injector) NewCommentService() service.ICommentService {
	return service.NewCommentService(i.NewCommentRepository())
}

func (i *injector) NewCommentHandler() handler.ICommentHandler {
	return handler.NewCommentHandler(i.NewCommentService())
}

// Post
func (i *injector) NewPostRepository() repository.IPostRepository {
	return persistence.NewPostRepository(i.conn)
}

func (i *injector) NewPostService() service.IPostService {
	return service.NewPostService(i.NewPostRepository())
}

func (i *injector) NewPostHandler() handler.IPostHandler {
	return handler.NewPostHandler(i.NewPostService())
}

// Tag
func (i *injector) NewTagRepository() repository.ITagRepository {
	return persistence.NewTagRepository(i.conn)
}

func (i *injector) NewTagService() service.ITagService {
	return service.NewTagService(i.NewTagRepository())
}

func (i *injector) NewTagHandler() handler.ITagHandler {
	return handler.NewTagHandler(i.NewTagService())
}

// Search
func (i *injector) NewSearchRepository() repository.ISearchRepository {
	return persistence.NewSearchRepository(i.conn)
}

func (i *injector) NewSearchService() service.ISearchService {
	return service.NewSearchService(i.NewSearchRepository())
}

func (i *injector) NewSearchHandler() handler.ISearchHandler {
	return handler.NewSearchHandler(i.NewSearchService())
}
