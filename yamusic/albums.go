package yamusic

import (
	"context"
	"fmt"
	"net/http"
)

type (
	// ArtistsService is a service to deal with artists.
	AlbumsService struct {
		client *Client
	}
	AlbumResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         Album          `json:"result"`
	}
)

// Get returns album
func (s *AlbumsService) Get(
	ctx context.Context,
	albumID int,
) (*AlbumResp, *http.Response, error) {
	if albumID == 0 {
		albumID = s.client.userID
	}

	uri := fmt.Sprintf("albums/%v", albumID)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	album := new(AlbumResp)
	resp, err := s.client.Do(ctx, req, album)
	return album, resp, err
}
