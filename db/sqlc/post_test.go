package db

import (
	"context"
	"testing"
	"time"

	"github.com/Kawaeugtkp/chepics_server/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func CreateRandomPost(t *testing.T) Post {
	user := CreateRandomUser(t)

	arg := CreatePostParams{
		OwnerID:  user.ID,
		Type:     util.RandomType(),
		Topic:    util.RandomString(20),
		Category: util.RandomCategory(),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.OwnerID, post.OwnerID)
	require.Equal(t, arg.Type, post.Type)
	require.Equal(t, arg.Topic, post.Topic)
	require.Equal(t, arg.Category, post.Category)

	require.Empty(t, post.IsRootOpinion)
	require.Empty(t, post.Description)
	require.Empty(t, post.Caption)
	require.Empty(t, post.TopicID)
	require.Empty(t, post.SetID)
	require.Empty(t, post.BaseOpinionID)
	require.Empty(t, post.PostImageUrl)
	require.Empty(t, post.Link)

	require.NotZero(t, post.ID)
	require.NotZero(t, post.Timestamp)

	return post
}

func TestCreatePost(t *testing.T) {
	CreateRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := CreateRandomPost(t)
	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.OwnerID, post2.OwnerID)
	require.Equal(t, post1.Type, post2.Type)
	require.Equal(t, post1.Topic, post2.Topic)
	require.Equal(t, post1.Category, post2.Category)
	require.Equal(t, post1.IsRootOpinion, post2.IsRootOpinion)
	require.Equal(t, post1.Description, post2.Description)
	require.Equal(t, post1.Caption, post2.Caption)
	require.Equal(t, post1.TopicID, post2.TopicID)
	require.Equal(t, post1.SetID, post2.SetID)
	require.Equal(t, post1.BaseOpinionID, post2.BaseOpinionID)
	require.Equal(t, post1.PostImageUrl, post2.PostImageUrl)
	require.Equal(t, post1.Link, post2.Link)
	require.Equal(t, post1.Votes, post2.Votes)

	require.WithinDuration(t, post1.Timestamp, post2.Timestamp, time.Second)
}

func TestUpdatePost(t *testing.T) {
	post1 := CreateRandomPost(t)

	arg := UpdatePostParams{
		ID:    post1.ID,
		Votes: post1.Votes + int32(1),
	}

	post2, err := testQueries.UpdatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.OwnerID, post2.OwnerID)
	require.Equal(t, post1.Type, post2.Type)
	require.Equal(t, post1.Topic, post2.Topic)
	require.Equal(t, post1.Category, post2.Category)
	require.Equal(t, post1.IsRootOpinion, post2.IsRootOpinion)
	require.Equal(t, post1.Description, post2.Description)
	require.Equal(t, post1.Caption, post2.Caption)
	require.Equal(t, post1.TopicID, post2.TopicID)
	require.Equal(t, post1.SetID, post2.SetID)
	require.Equal(t, post1.BaseOpinionID, post2.BaseOpinionID)
	require.Equal(t, post1.PostImageUrl, post2.PostImageUrl)
	require.Equal(t, post1.Link, post2.Link)
	require.Equal(t, arg.Votes, post2.Votes)
	require.Equal(t, post1.Timestamp, post2.Timestamp)
}

func TestDeletePost(t *testing.T) {
	post1 := CreateRandomPost(t)
	err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)

	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, post2)
}

func TestListPosts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomPost(t)
	}

	arg := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}

	posts, err := testQueries.ListPosts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, posts, 5)

	for _, post := range posts {
		require.NotEmpty(t, post)
	}
}
