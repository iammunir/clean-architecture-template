package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/iammunir/clean-architecture-template/app"
	"github.com/iammunir/clean-architecture-template/database"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	mysqlConn, errMysql := database.InitMySQL()
	if errMysql != nil {
		log.Fatal("Error connection with MySQL: ", errMysql.Error())
	}

	router := app.InitRouter(mysqlConn)
	log.Println("routes Initialized")

	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Println("Server Initialized at port: ", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
