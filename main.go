package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	DB "newCurriculum/db"
	"newCurriculum/gql"
	"newCurriculum/gql/resolver"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8000"

func main() {
	fmt.Println("Hello World")

	// DBに接続
	POSTGRES_USER := os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB := os.Getenv("POSTGRES_DB")
	if os.Getenv("USE_TEST_DB") == "true" {
		POSTGRES_DB = os.Getenv("POSTGRES_DB")
		fmt.Println("[notice] this is a test")
	}
	connStr := "host=postgres port=5432 user=" + POSTGRES_USER + " password=" + POSTGRES_PASSWORD + " dbname=" + POSTGRES_DB + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetDB(db)
	// ポートでlisten
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &resolver.Resolver{}}))
	// errorをここで、一括で出力
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		log.Print(err)
		return err
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
