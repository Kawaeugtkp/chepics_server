package api

import (
	"database/sql"
	"net/http"

	db "github.com/Kawaeugtkp/chepics_server/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	OwnerID       int64  `json:"owner_id" binding:"required"`
	Type          string `json:"type" binding:"required,oneof=topic opinion"`
	IsRootOpinion *bool   `json:"is_root_opinion"`
	Topic         string `json:"topic" binding:"required"`
	Description   *string `json:"description"`
	Caption       *string `json:"caption"`
	TopicID       *int64  `json:"topic_id"`
	SetID         *int64  `json:"set_id"`
	Category      string `json:"category" binding:"required,oneof=news sport entertainment covid economy tech fashion life gourmet browse culture anime funny love"`
	BaseOpinionID *int64  `json:"base_opinion_id"`
	PostImageUrl  *string `json:"post_image_url"`
	Link          *string `json:"link"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		OwnerID:       req.OwnerID,
		Type:          req.Type,
		IsRootOpinion: req.IsRootOpinion,
		Topic:         req.Topic,
		Description:   req.Description,
		Caption:       req.Caption,
		TopicID:       req.TopicID,
		SetID:         req.SetID,
		Category:      req.Category,
		BaseOpinionID: req.BaseOpinionID,
		PostImageUrl:  req.PostImageUrl,
		Link:          req.Link,
	}

	post, err := server.store.CreatePost(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type getPostRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.store.GetPost(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type listPostRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) listPost(ctx *gin.Context) {
	var req listPostRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPostsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	posts, err := server.store.ListPosts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}