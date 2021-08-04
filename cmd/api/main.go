package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/antonlindstrom/pgstore"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/api/handler"
	persistence "github.com/oshou/AwesomeMusic-api/api/infra/persistence/postgres"
	mw "github.com/oshou/AwesomeMusic-api/api/middleware"
	"github.com/oshou/AwesomeMusic-api/api/usecase"
	"github.com/oshou/AwesomeMusic-api/config"
	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/log"
	"github.com/oshou/AwesomeMusic-api/sentry"
)

const (
	name    = "api"
	version = "v1.5"

	httpGzipLevel        = 6
	httpTimeoutSecond    = 60
	httpPortString       = ":8080"
	corsMaxAgeSecond     = 300
	filePath             = "./config/config.yml"
	sessionClearInterval = 5
)

var (
	env  string = "local"
	port string = "8080"
)

func main() {
	flag.StringVar(&port, "port", httpPortString, "tcp host:port to connect")

	// Logger
	log.Init()
	// nolint: errcheck
	defer log.Logger.Sync()

	// Config
	conf, err := config.NewConfig(filePath)
	if err != nil {
		log.Logger.Fatal("failed to initialize config", zap.Error(err))
	}

	// DBConnection
	db, err := db.NewDB(conf)
	if err != nil {
		log.Logger.Fatal("failed to initialize db", zap.Error(err))
	}
	defer db.Close()

	// Session Store
	dsn := conf.GetDSN()
	store, err := pgstore.NewPGStore(dsn, []byte(os.Getenv("SESSION_SECRET_KEY")))
	if err != nil {
		log.Logger.Fatal("failed to initialize session store", zap.Error(err))
	}

	store.Options = &sessions.Options{
		Path:     "/",
		Domain:   os.Getenv("COOKIE_DOMAIN"),
		MaxAge:   60 * 60 * 1,
		Secure:   false,
		HttpOnly: true,
	}
	defer store.Close()
	defer store.StopCleanup(store.Cleanup(time.Duration(sessionClearInterval) * time.Minute))

	// Error Reporter
	sentryDSN := conf.GetSentryDSN()
	if sentryDSN == "" {
		log.Logger.Warn("not setup sentry", zap.Error(err))
	} else {
		if err = sentry.Init(sentryDSN, os.Getenv("GO_ENV"), name, version); err != nil {
			log.Logger.Warn("failed to initialize sentry", zap.Error(err))
		}
	}
	defer sentry.Recover()
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})

	// Injector
	healthRepository := persistence.NewHealthRepository(db.Pool)
	userRepository := persistence.NewUserRepository(db.Pool)
	commentRepository := persistence.NewCommentRepository(db.Pool)
	postRepository := persistence.NewPostRepository(db.Pool)
	tagRepository := persistence.NewTagRepository(db.Pool)
	searchRepository := persistence.NewSearchRepository(db.Pool)

	healthUsecase := usecase.NewHealthUsecase(healthRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)
	commentUsecase := usecase.NewCommentUsecase(commentRepository)
	postUsecase := usecase.NewPostUsecase(postRepository)
	tagUsecase := usecase.NewTagUsecase(tagRepository)
	searchUsecase := usecase.NewSearchUsecase(searchRepository)

	healthHandler := handler.NewHealthHandler(healthUsecase)
	authHandler := handler.NewAuthHandler(userUsecase, store)
	userHandler := handler.NewUserHandler(userUsecase)
	commentHandler := handler.NewCommentHandler(commentUsecase)
	postHandler := handler.NewPostHandler(postUsecase)
	tagHandler := handler.NewTagHandler(tagUsecase)
	searchHandler := handler.NewSearchHandler(searchUsecase)

	// Router
	r := chi.NewRouter()

	// Middlware
	r.Use(
		mw.ZapLogger(log.Logger),
		middleware.RequestID,
		middleware.StripSlashes,
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
		sentryHandler.Handle,
	)

	// Routing
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", healthHandler.GetHealth)
		r.Post("/login", authHandler.Login)
		r.Post("/logout", authHandler.Logout)
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.ListUsers)
			r.Post("/", userHandler.AddUser)
			r.Get("/{user_id}", userHandler.GetUserByID)
		})
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", postHandler.ListPosts)
			r.Post("/", postHandler.AddPost)
			r.Route("/{post_id}", func(r chi.Router) {
				r.Get("/", postHandler.GetPostByID)
				r.Delete("/", postHandler.DeletePostByID)
				r.Route("/comments", func(r chi.Router) {
					r.Get("/", commentHandler.ListComments)
					r.Post("/", commentHandler.AddComment)
					r.Get("/{comment_id}", commentHandler.GetCommentByID)
				})
				r.Route("/tags", func(r chi.Router) {
					r.Get("/", tagHandler.ListTagsByPostID)
					r.Post("/{tag_id}", tagHandler.AttachTag)
				})
			})
		})
		r.Route("/tags", func(r chi.Router) {
			r.Get("/", tagHandler.ListTags)
			r.Post("/", tagHandler.AddTag)
			r.Get("/{tag_id}", tagHandler.GetTagByID)
		})
		r.Route("/search", func(r chi.Router) {
			r.Get("/", searchHandler.SearchByType)
		})
	})

	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Logger.Error("failed to listen server", zap.Error(err))
			os.Exit(1)
		}
	}()

	log.Logger.Info("start server")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Logger.Error("err", zap.Error(err))
	}

	log.Logger.Info("shutdown server")
}
