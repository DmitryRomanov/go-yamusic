//go:build integration

package integration

import (
	"context"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArtistAlbums(t *testing.T) {
	var artistID int = 9462344
	ctx := context.Background()
	t.Run("Get artist's albums", func(t *testing.T) {
		result, resp, err := client.Artists().GetDirectAlbums(ctx, artistID)
		require.NoError(t, err)
		require.NotZero(t, result)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.NotZero(t, result.Result)
		require.Equal(t, result.Result.Albums[0].Artists[0].ID, artistID)
	})
}

func TestArtistInfo(t *testing.T) {
	var artistID int = 9462344
	ctx := context.Background()
	t.Run("Get artist's albums", func(t *testing.T) {
		result, resp, err := client.Artists().GetBriefInfo(ctx, artistID)
		require.NoError(t, err)
		require.NotZero(t, result)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.NotZero(t, result.Result)
		require.Equal(t, result.Result.Artist.ID, strconv.Itoa(artistID))
	})
}
