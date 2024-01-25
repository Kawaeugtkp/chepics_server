package api

import (
	db "github.com/Kawaeugtkp/chepics_server/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/vanng822/go-solr/solr"
)

// Server serves HTTP requests for chepics service.
type Server struct {
	store *db.Store
	solrInterface *solr.SolrInterface
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store, solrInterface *solr.SolrInterface) *Server {
	server := &Server{
		store: store,
		solrInterface: solrInterface,
	}
	router := gin.Default()

	router.POST("/posts", server.createPost)
	router.GET("/posts/:id", server.getPost)
	router.GET("/posts", server.listPost)
	router.GET("/search/posts", server.listSolrPost)

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