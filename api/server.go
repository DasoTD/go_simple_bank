package api

import (
	"fmt"
	"io"
	"os"
	"time"

	db "github.com/dasotd/go_simple_bank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	 // Logging to a file.
	 f, _ := os.Create("gin.log")
	 gin.DefaultWriter = io.MultiWriter(f)

	 
	router  := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}


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

	router.POST("/users", server.createUser)


	router.POST("/accounts", server.createAccount)
	router.GET("/account/:id", server.getAccount)
	router.GET("/account", server.listAccounts)
	router.DELETE("/account/:id", server.deleteAccount)


	// Entries Router
	router.POST("/entry", server.createEntry)
	router.GET("/entry/:id", server.getEntry)
	router.GET("/entries", server.ListEntry)
	router.DELETE("/entry/:id", server.deleteEntry)

	router.POST("/transfer", server.createTransfer)


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