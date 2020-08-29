package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	//_ "github.com/go-sql-driver/mysql"

	"github.com/oshou/AwesomeMusic-api/db"
	persistence "github.com/oshou/AwesomeMusic-api/infrastructure/persistence/postgres"
	"github.com/oshou/AwesomeMusic-api/log"
	"github.com/oshou/AwesomeMusic-api/ui/http/handler"
	"github.com/oshou/AwesomeMusic-api/ui/http/session"

	//mw "github.com/oshou/AwesomeMusic-api/ui/http/middleware"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

const (
	httpGzipLevel     = 6
	httpTimeoutSecond = 60
	httpPortString    = ":8080"
	corsMaxAgeSecond  = 300
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
	if err := db.Init(); err != nil {
		log.Logger.Fatal("failed to connect db", zap.Error(err))
	}
	defer db.Close()
	pool := db.GetDB()

	// Session
	sskey := os.Getenv("SESSION_SECRET_KEY")
	opt := &sessions.Options{
		Path:     "/",
		Domain:   os.Getenv("COOKIE_DOMAIN"),
		MaxAge:   60 * 60 * 1,
		Secure:   false,
		HttpOnly: true,
	}
	ssstore, err := session.NewStore(sskey, opt)

	if err != nil {
		log.Logger.Fatal("failed to initialize session store", zap.Error(err))
	}

	// Injector
	healthRepository := persistence.NewHealthRepository(pool)
	userRepository := persistence.NewUserRepository(pool)
	commentRepository := persistence.NewCommentRepository(pool)
	postRepository := persistence.NewPostRepository(pool)
	tagRepository := persistence.NewTagRepository(pool)
	searchRepository := persistence.NewSearchRepository(pool)

	healthUsecase := usecase.NewHealthUsecase(healthRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)
	postUsecase := usecase.NewPostUsecase(postRepository)
	tagUsecase := usecase.NewTagUsecase(tagRepository)
	searchUsecase := usecase.NewSearchUsecase(searchRepository)

	healthHandler := handler.NewHealthHandler(healthUsecase)
	loginHandler := handler.NewLoginHandler(userUsecase, ssstore)
	userHandler := handler.NewUserHandler(userUsecase)
	commentHandler := handler.NewCommentHandler(commentUsecase)
	postHandler := handler.NewPostHandler(postUsecase)
	tagHandler := handler.NewTagHandler(tagUsecase)
	searchHandler := handler.NewSearchHandler(searchUsecase)

	// Router
	r := chi.NewRouter()

	// Middlware
	r.Use(
		middleware.Logger,
		//mw.ZapLogger(log.Logger),
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
		r.Get("/health", healthHandler.Health)
		r.Post("/login", loginHandler.Login)
		r.Post("/logout", loginHandler.Logout)
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
					r.Get("/", tagHandler.GetTagsByPostID)
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

	log.Logger.Info("start server")
	if err := srv.ListenAndServe(); err != nil {
		log.Logger.Fatal("failed to start server", zap.Error(err))
	}
}
