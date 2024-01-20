// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	DeletePost(ctx context.Context, id int64) error
	GetPost(ctx context.Context, id int64) (Post, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
}

var _ Querier = (*Queries)(nil)
