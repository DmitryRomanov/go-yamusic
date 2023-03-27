package yamusic

import (
	"context"
	"fmt"
	"net/http"
)

type (
	// ArtistsService is a service to deal with artists.
	ArtistsService struct {
		client *Client
	}

	// AlbumsResp describes get artists's albums response
	AlbumsResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         struct {
			Pager struct {
				Total   int `json:"total"`
				Page    int `json:"page"`
				PerPage int `json:"perPage"`
			} `json:"pager"`
			Albums []Album `json:"albums"`
		} `json:"result"`
	}
)

// GetDirectAlbums returns artists's albums
func (s *ArtistsService) GetDirectAlbums(
	ctx context.Context,
	artistID int,
) (*AlbumsResp, *http.Response, error) {
	if artistID == 0 {
		artistID = s.client.userID
	}

	uri := fmt.Sprintf("artists/%v/direct-albums", artistID)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	albums := new(AlbumsResp)
	resp, err := s.client.Do(ctx, req, albums)
	return albums, resp, err
}
