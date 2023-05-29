package api

import (
	"fmt"
	"io"
	"os"
	"time"

	db "github.com/dasotd/go_simple_bank/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	 // Logging to a file.
	 f, _ := os.Create("gin.log")
	 gin.DefaultWriter = io.MultiWriter(f)

	 
	router  := gin.Default()
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

	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccounts)
	router.DELETE("/account/:id", server.deleteAccount)


	server.router = router
	return server

}

func errorResponse(err error) gin.H {
		return gin.H{"error": err.Error()}
	}

	// // // Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}