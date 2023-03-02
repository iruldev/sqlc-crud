package api

import (
	"github.com/gin-gonic/gin"
	database "github.com/iruldev/sqlc-crud/database/sqlc"
)

// Server serves HTTP requests for our booking service
type Server struct {
	store  database.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store database.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
