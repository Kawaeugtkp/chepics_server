package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vanng822/go-solr/solr"
)

type listSolrPostRequest struct {
	// 大文字！！！！！！
	Key  string `form:"key" binding:"required"`
	Word string `form:"word" binding:"required"`
}

func (server *Server) listSolrPost(ctx *gin.Context) {
	var req listSolrPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	query := solr.NewQuery()
	fmt.Println("DEBUG: key is ", req.Key)
	fmt.Println("DEBUG: word is ", req.Word)
	queryString := fmt.Sprintf("%s:%s", req.Key, req.Word)
	fmt.Println("DEBUG: ", queryString)
	query.Q(queryString)
	search := server.solrInterface.Search(query)
	result, err := search.Result(nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result.Results.Docs)
}
