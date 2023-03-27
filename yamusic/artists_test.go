package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtistsService_GetDirectAlbums(t *testing.T) {
	setup()
	defer teardown()

	want := &AlbumsResp{}
	want.InvocationInfo.ReqID = "Artists.DirectAlbums"
	artistID := 100

	mux.HandleFunc(
		fmt.Sprintf("/artists/%v/direct-albums", artistID),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Artists().GetDirectAlbums(context.Background(), artistID)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
