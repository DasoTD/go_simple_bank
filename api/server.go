package api

import (
	"fmt"
	"io"
	"os"
	_"time"

	db "github.com/dasotd/go_simple_bank/db/sqlc"
	"github.com/dasotd/go_simple_bank/token"
	"github.com/dasotd/go_simple_bank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	 // Logging to a file.
	 f, _ := os.Create("gin.log")
	 gin.DefaultWriter = io.MultiWriter(f)

	 
	

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}


	// router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

	// 	// your custom format
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	//   }))

	


	


	server.setupRouter()
	return server, nil
	}

	func (server *Server) setupRouter() {
			router  := gin.Default()
			router.POST("/users", server.createUser)
			router.POST("/users/login", server.loginUser)
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
	}



func errorResponse(err error) gin.H {
		return gin.H{"error": err.Error()}
	}

	// // // Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}