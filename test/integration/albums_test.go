//go:build integration

package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlbums(t *testing.T) {
	var kind int = 9911104
	ctx := context.Background()
	t.Run("Get album information", func(t *testing.T) {
		result, resp, err := client.Albums().Get(ctx, kind)
		require.NoError(t, err)
		require.NotZero(t, result)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.NotZero(t, result.Result)
		require.Equal(t, result.Result.ID, kind)
	})
}
