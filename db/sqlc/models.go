// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"
)

type Pickset struct {
	ID      int64  `json:"id"`
	TopicID int64  `json:"topic_id"`
	Title   string `json:"title"`
	Votes   int32  `json:"votes"`
}

type Post struct {
	ID            int64          `json:"id"`
	Timestamp     time.Time      `json:"timestamp"`
	OwnerID       int64          `json:"owner_id"`
	Type          string         `json:"type"`
	IsRootOpinion sql.NullBool   `json:"is_root_opinion"`
	Votes         int32          `json:"votes"`
	Topic         string         `json:"topic"`
	Description   sql.NullString `json:"description"`
	Caption       sql.NullString `json:"caption"`
	TopicID       sql.NullInt64  `json:"topic_id"`
	SetID         sql.NullInt64  `json:"set_id"`
	Category      string         `json:"category"`
	BaseOpinionID sql.NullInt64  `json:"base_opinion_id"`
	PostImageUrl  sql.NullString `json:"post_image_url"`
	Link          sql.NullString `json:"link"`
}

type User struct {
	ID              int64  `json:"id"`
	Username        string `json:"username"`
	HashedPassword  string `json:"hashed_password"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"profile_image_url"`
	Bio             string `json:"bio"`
	Follower        int32  `json:"follower"`
	Following       int32  `json:"following"`
}
