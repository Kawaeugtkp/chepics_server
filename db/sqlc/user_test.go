package db

import (
	"context"
	"testing"

	"github.com/Kawaeugtkp/chepics_server/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:        util.RandomUsername(),
		HashedPassword:  util.RandomString(6),
		FullName:        util.RandomString(6),
		Email:           util.RandomEmail(),
		ProfileImageUrl: util.RandomString(6),
		Bio:             util.RandomString(6),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.ProfileImageUrl, user.ProfileImageUrl)
	require.Equal(t, arg.Bio, user.Bio)
	require.Equal(t, arg.Follower, user.Follower)
	require.Equal(t, arg.Following, user.Following)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.ProfileImageUrl, user2.ProfileImageUrl)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.Follower, user2.Follower)
	require.Equal(t, user1.Following, user2.Following)
	require.Equal(t, user2.Follower, int32(0))
	require.Equal(t, user2.Following, int32(0))
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)

	arg := UpdateUserParams{
		ID:              user1.ID,
		Username:        util.RandomUsername(),
		FullName:        util.RandomString(6),
		ProfileImageUrl: util.RandomString(6),
		Bio:             util.RandomString(20),
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.ProfileImageUrl, user2.ProfileImageUrl)
	require.Equal(t, arg.Bio, user2.Bio)
	require.Equal(t, user1.Follower, user2.Follower)
	require.Equal(t, user1.Following, user2.Following)
}

func TestDeleteUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
