package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type listSolrPostRequest struct {
	// 大文字！！！！！！
	Word string `form:"word" binding:"required"`
}

func (server *Server) listSolrPost(ctx *gin.Context) {
	var req listSolrPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	encodedQuery := url.QueryEscape(req.Word)
	url := fmt.Sprintf("http://localhost:8984/solr/post/search?indent=true&start=0&q=%s", encodedQuery)

	resp, err := http.Get(url)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := result["response"].(map[string]interface{})
	docs := response["docs"].([]interface{})

	ctx.JSON(http.StatusOK, docs)
}
