package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	//_ "github.com/go-sql-driver/mysql"

	"github.com/oshou/AwesomeMusic-api/db"
	persistence "github.com/oshou/AwesomeMusic-api/infrastructure/persistence/postgres"
	"github.com/oshou/AwesomeMusic-api/log"
	"github.com/oshou/AwesomeMusic-api/ui/http/handler"
	mw "github.com/oshou/AwesomeMusic-api/ui/http/middleware"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

const (
	httpGzipLevel     = 6
	corsMaxAgeSecond  = 300
	httpTimeoutSecond = 60
	httpPortString    = ":8080"
)

func main() {
	// Set Logger
	log.Init()
	defer log.Logger.Sync()

	// Load Environment
	if err := godotenv.Load(); err != nil {
		log.Logger.Fatal("failed to loading .env file", zap.Error(err))
	}

	// DB Connection
	err := db.Init()
	if err != nil {
		log.Logger.Fatal("failed to connect db", zap.Error(err))
	}
	defer func() {
		err := db.Pool.Close()
		if err != nil {
			log.Logger.Fatal("failed to release db", zap.Error(err))
		}
	}()

	// Injector
	userRepository := persistence.NewUserRepository(db.Pool)
	commentRepository := persistence.NewCommentRepository(db.Pool)
	postRepository := persistence.NewPostRepository(db.Pool)
	tagRepository := persistence.NewTagRepository(db.Pool)
	searchRepository := persistence.NewSearchRepository(db.Pool)

	userUsecase := usecase.NewUserUsecase(userRepository)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)
	postUsecase := usecase.NewPostUsecase(postRepository)
	tagUsecase := usecase.NewTagUsecase(tagRepository)
	searchUsecase := usecase.NewSearchUsecase(searchRepository)

	userHandler := handler.NewUserHandler(userUsecase)
	commentHandler := handler.NewCommentHandler(commentUsecase)
	postHandler := handler.NewPostHandler(postUsecase)
	tagHandler := handler.NewTagHandler(tagUsecase)
	searchHandler := handler.NewSearchHandler(searchUsecase)

	r := chi.NewRouter()

	// Middlware
	r.Use(
		//middleware.Logger,
		mw.ZapLogger(log.Logger),
		middleware.Recoverer,
		middleware.SetHeader("Content-Type", "application/json"),
		middleware.Compress(httpGzipLevel, "gzip"),
		middleware.Timeout(httpTimeoutSecond),
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           corsMaxAgeSecond,
		}),
	)

	// Routing
	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.GetUsers)
			r.Post("/", userHandler.AddUser)
			r.Get("/{user_id}", userHandler.GetUserByID)
		})
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", postHandler.GetPosts)
			r.Post("/", postHandler.AddPost)
			r.Route("/{post_id}", func(r chi.Router) {
				r.Get("/", postHandler.GetPostByID)
				r.Delete("/", postHandler.DeletePostByID)
				r.Route("/comments", func(r chi.Router) {
					r.Get("/", commentHandler.GetComments)
					r.Post("/", commentHandler.AddComment)
					r.Get("/{comment_id}", commentHandler.GetCommentByID)
				})
				r.Route("/tags", func(r chi.Router) {
					r.Post("/", tagHandler.GetTagsByPostID)
					r.Post("/{tag_id}", tagHandler.AttachTag)
				})
			})
		})
		r.Route("/tags", func(r chi.Router) {
			r.Get("/", tagHandler.GetTags)
			r.Post("/", tagHandler.AddTag)
			r.Get("/{tag_id}", tagHandler.GetTagByID)
		})
		r.Route("/search", func(r chi.Router) {
			r.Get("/", searchHandler.SearchByType)
		})
	})

	srv := &http.Server{
		Addr:    httpPortString,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Logger.Fatal("failed to start server", zap.Error(err))
	}
}
