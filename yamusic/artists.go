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

	// ArtistResp describes get artist response
	ArtistResp struct {
		InvocationInfo InvocationInfo `json:"invocationInfo"`
		Error          Error          `json:"error"`
		Result         struct {
			Artist struct {
				ID               string   `json:"id"`
				Composer         bool     `json:"composer"`
				Various          bool     `json:"various"`
				TicketsAvailable bool     `json:"ticketsAvailable"`
				Name             string   `json:"name"`
				Genres           []string `json:"genres"`
				Regions          []string `json:"regions"`
				Cover            struct {
					Type   string `json:"type"`
					Prefix string `json:"prefix"`
					URI    string `json:"uri"`
				} `json:"cover"`
				Counts struct {
					Tracks       int `json:"tracks"`
					DirectAlbums int `json:"directAlbums"`
					AlsoAlbums   int `json:"alsoAlbums"`
					AlsoTracks   int `json:"alsoTracks"`
				} `json:"counts"`
				Ratings struct {
					Day   int `json:"day"`
					Week  int `json:"week"`
					Month int `json:"month"`
				} `json:"ratings,omitempty"`
				Links []struct {
					Title         string `json:"title"`
					Href          string `json:"href"`
					Type          string `json:"type"`
					SocialNetwork string `json:"socialNetwork,omitempty"`
				} `json:"links"`
			} `json:"artist"`
			Albums Albums `json:"albums"`
			Stats  struct {
				LastMonthListeners int `json:"lastMonthListeners"`
			} `json:"stats"`
			SimilarArtist Artists `json:"similarArtists"`
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

// GetBriefInfo returns artist info
func (s *ArtistsService) GetBriefInfo(
	ctx context.Context,
	artistID int,
) (*ArtistResp, *http.Response, error) {
	if artistID == 0 {
		artistID = s.client.userID
	}

	uri := fmt.Sprintf("artists/%v/brief-info", artistID)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	artist := new(ArtistResp)
	resp, err := s.client.Do(ctx, req, artist)

	return artist, resp, err
}
