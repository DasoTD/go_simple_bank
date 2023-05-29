package main

import (
	"database/sql"
	"fmt"
	"log"
	_"net/http"
	"time"
	_"github.com/lib/pq"
	"github.com/dasotd/go_simple_bank/api"
	db "github.com/dasotd/go_simple_bank/sqlc"
	"github.com/gin-gonic/gin"
)


const (
	DBDriver ="postgres"
	DBSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress= "0.0.0.0:8080"
)


func main(){

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	  }))

	con, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(con)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}

// 	router.Use(gin.Recovery())

//   router.GET("/ping", func(c *gin.Context) {
//     c.String(http.StatusOK, "pong")
//   })

// 	router.Run(":8080")
	
}