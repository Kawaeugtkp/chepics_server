package api

import (
	db "github.com/Kawaeugtkp/chepics_server/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for chepics service.
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/posts", server.createPost)
	router.GET("/posts/:id", server.getPost)
	router.GET("/posts", server.listPost)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}