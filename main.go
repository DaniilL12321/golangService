package main

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"golangService/internal/application"
	"golangService/internal/repository"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	dbpool, err := repository.InitDBConn(ctx)
	if err != nil {
		log.Fatalf("%w failed to init DB connection", err)
	}
	defer dbpool.Close()
	a := application.NewApp(ctx, dbpool)
	r := httprouter.New()
	a.Routes(r)
	srv := &http.Server{Addr: ":8080", Handler: r}
	fmt.Println("Server running on port 8080")
	srv.ListenAndServe()
}
