package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "net/http"
	"time"

	"github.com/dasotd/go_simple_bank/api"
	db "github.com/dasotd/go_simple_bank/db/sqlc"
	"github.com/dasotd/go_simple_bank/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)





func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config")
	}
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

	con, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(con)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("can not start server", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server", err)
	}

// 	router.Use(gin.Recovery())

//   router.GET("/ping", func(c *gin.Context) {
//     c.String(http.StatusOK, "pong")
//   })

// 	router.Run(":8080")
	
}