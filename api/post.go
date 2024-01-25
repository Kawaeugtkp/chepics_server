package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/Kawaeugtkp/chepics_server/db/sqlc"
	"github.com/Kawaeugtkp/chepics_server/util"
	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	OwnerID       int64  `json:"owner_id" binding:"required"`
	Type          string `json:"type" binding:"required,oneof=topic opinion"`
	IsRootOpinion bool   `json:"is_root_opinion"`
	Topic         string `json:"topic" binding:"required"`
	Description   string `json:"description"`
	Caption       string `json:"caption"`
	TopicID       int64  `json:"topic_id"`
	SetID         int64  `json:"set_id"`
	Category      string `json:"category" binding:"required,oneof=news sport entertainment covid economy tech fashion life gourmet browse culture anime funny love"`
	BaseOpinionID int64  `json:"base_opinion_id"`
	PostImageUrl  string `json:"post_image_url"`
	Link          string `json:"link"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// TODO: - 自前判断は流石にやばいので解決法を聞く
	arg := db.CreatePostParams{
		OwnerID:       req.OwnerID,
		Type:          req.Type,
		IsRootOpinion: sql.NullBool{Bool: req.IsRootOpinion},
		Topic:         req.Topic,
		Description:   sql.NullString{String: req.Description, Valid: util.CheckNonNullString(req.Description)},
		Caption:       sql.NullString{String: req.Caption, Valid: util.CheckNonNullString(req.Caption)},
		TopicID:       sql.NullInt64{Int64: req.TopicID, Valid: util.CheckNonNullInt64(req.TopicID)},
		SetID:         sql.NullInt64{Int64: req.SetID, Valid: util.CheckNonNullInt64(req.SetID)},
		Category:      req.Category,
		BaseOpinionID: sql.NullInt64{Int64: req.BaseOpinionID, Valid: util.CheckNonNullInt64(req.BaseOpinionID)},
		PostImageUrl:  sql.NullString{String: req.PostImageUrl, Valid: util.CheckNonNullString(req.PostImageUrl)},
		Link:          sql.NullString{String: req.Link, Valid: util.CheckNonNullString(req.Link)},
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

	ctx.JSON(http.StatusOK, getPostResponse(post))
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

	var responses []PostResponse

	for _, post := range posts {
		responses = append(responses, getPostResponse(post))
	}

	ctx.JSON(http.StatusOK, responses)
}

type PostResponse struct {
	ID            int64     `json:"id"`
	Timestamp     time.Time `json:"timestamp"`
	OwnerID       int64     `json:"owner_id"`
	Type          string    `json:"type"`
	IsRootOpinion *bool     `json:"is_root_opinion"`
	Votes         int32     `json:"votes"`
	Topic         string    `json:"topic"`
	Description   *string   `json:"description"`
	Caption       *string   `json:"caption"`
	TopicID       *int64    `json:"topic_id"`
	SetID         *int64    `json:"set_id"`
	Category      string    `json:"category"`
	BaseOpinionID *int64    `json:"base_opinion_id"`
	PostImageUrl  *string   `json:"post_image_url"`
	Link          *string   `json:"link"`
}

func getPostResponse(post db.Post) PostResponse {
	var isRootOpinion *bool
	var description *string
	var caption *string
	var topicId *int64
	var setId *int64
	var baseOpinionId *int64
	var postImageUrl *string
	var link *string

	if post.IsRootOpinion.Valid {
		isRootOpinion = &post.IsRootOpinion.Bool
	}

	if post.Description.Valid {
		description = &post.Description.String
	}

	if post.Caption.Valid {
		caption = &post.Caption.String
	}

	if post.TopicID.Valid {
		topicId = &post.TopicID.Int64
	}

	if post.SetID.Valid {
		setId = &post.SetID.Int64
	}

	if post.BaseOpinionID.Valid {
		baseOpinionId = &post.BaseOpinionID.Int64
	}

	if post.PostImageUrl.Valid {
		postImageUrl = &post.PostImageUrl.String
	}

	if post.Link.Valid {
		link = &post.Link.String
	}

	response := PostResponse{
		ID: post.ID,
		Timestamp: post.Timestamp,
		OwnerID: post.OwnerID,
		Type: post.Type,
		IsRootOpinion: isRootOpinion,
		Votes: post.Votes,
		Topic: post.Topic,
		Description: description,
		Caption: caption,
		TopicID: topicId,
		SetID: setId,
		Category: post.Category,
		BaseOpinionID: baseOpinionId,
		PostImageUrl: postImageUrl,
		Link: link,
	}

	return response
}
