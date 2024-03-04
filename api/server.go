package api

import (
	"fmt"

	db "github.com/BinhNguyenDang/simplebank/db/sqlc"
	"github.com/BinhNguyenDang/simplebank/token"
	"github.com/BinhNguyenDang/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding" //sub package of gin
	"github.com/go-playground/validator/v10"
)

// server serves HTTP requests for our banking service.
type Server struct {
	config util.Config
	store  db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %d", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	// to get the current validator engine that gin using 
	// Engine() return  a general interface type which by default is a pointer to validator onjects
	// convert output to a validator pointer
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency", validCurrency)
	}
	

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter(){
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PUT("/accounts/:id",server.updateAccount)

	router.POST("/transfers", server.createTransfer)
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
