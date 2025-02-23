package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlbumsSevice_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &TrackResp{}
	want.InvocationInfo.ReqID = "Albums.Get"

	kind := 42

	mux.HandleFunc(
		fmt.Sprintf("/albums/%d", kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Albums().Get(
		context.Background(),
		kind,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
