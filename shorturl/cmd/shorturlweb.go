package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"shorten-url-go/shorturl/healthchecker"
	"shorten-url-go/shorturl/inmem"
	"shorten-url-go/shorturl/managing"
	"shorten-url-go/shorturl/parser"
	"shorten-url-go/shorturl/redirecting"
)

func main() {
	repo := inmem.NewShortURLRepository()
	parserer := parser.NewParser()

	managingsrv := managing.NewService(repo)
	managingTransport := managing.NewTransport(managingsrv, parserer)

	listingsrv := redirecting.NewService(repo)
	listingTransport := redirecting.NewTransport(listingsrv)

	ctx := context.Background()
	receiver := healthchecker.NewHealthChecker(time.Minute*5, repo)
	go receiver.Start(ctx)

	http.HandleFunc("/", listingTransport.HandleRedirect)
	http.HandleFunc("/shorturl", managingTransport.HandleManage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
