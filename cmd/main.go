package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	//_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"

	"github.com/oshou/AwesomeMusic-api/db"
	persistence "github.com/oshou/AwesomeMusic-api/infrastructure/persistence/postgres"
	"github.com/oshou/AwesomeMusic-api/ui/http/handler"
	"github.com/oshou/AwesomeMusic-api/usecase"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	// Load Environment
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading .env file: %+v\n", err)
		log.Fatalln()
	}

	// Set DBConnection
	db, err := db.NewDB()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}()

	// Routing
	//i := injector.NewInjector(db)
	//h := i.NewAppHandler()
	//r := i.NewRouter(h)

	r := chi.NewRouter()

	// Injector
	userRepository := persistence.NewUserRepository(db)
	commentRepository := persistence.NewCommentRepository(db)
	postRepository := persistence.NewPostRepository(db)
	tagRepository := persistence.NewTagRepository(db)
	searchRepository := persistence.NewSearchRepository(db)

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
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Print("err1 port:", os.Getenv("API_SERVER_PORT"), " ", err)
			os.Exit(1)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("err2")
	}
}
