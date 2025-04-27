package lavalyrics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/disgoorg/disgolink/v4/disgolink"
	"github.com/disgoorg/disgolink/v4/lavalink"
	"github.com/disgoorg/json/v2"
	"github.com/disgoorg/snowflake/v2"
)

type Lyrics struct {
	SourceName string           `json:"sourceName"`
	Provider   string           `json:"provider"`
	Text       string           `json:"text"`
	Lines      []Line           `json:"lines"`
	Plugin     lavalink.RawData `json:"plugin"`
}

type Line struct {
	Timestamp lavalink.Duration `json:"timestamp"`
	Duration  lavalink.Duration `json:"duration"`
	Line      string            `json:"line"`
	Plugin    lavalink.RawData  `json:"plugin"`
}

// GetCurrentTrackLyrics returns the lyrics of the current track being played in the guild.
// If the current track has no lyrics, it will return nil.
func GetCurrentTrackLyrics(ctx context.Context, client disgolink.RestClient, sessionID string, guildID snowflake.ID, skipTrackSource bool) (*Lyrics, error) {
	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/v4/sessions/%s/players/%s/track/lyrics", sessionID, guildID), nil)
	if err != nil {
		return nil, err
	}

	q := rq.URL.Query()
	q.Add("skipTrackSource", fmt.Sprint(skipTrackSource))
	rq.URL.RawQuery = q.Encode()

	rs, err := client.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()

	if rs.StatusCode < 200 || rs.StatusCode >= 300 {
		var lavalinkError lavalink.Error
		if err = json.NewDecoder(rs.Body).Decode(&lavalinkError); err != nil {
			return nil, err
		}
		return nil, lavalinkError
	}

	if rs.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	var l Lyrics
	if err = json.NewDecoder(rs.Body).Decode(&l); err != nil {
		return nil, err
	}
	return &l, nil
}

// GetLyrics returns the lyrics of the provided track.
// If the track has no lyrics, it will return nil.
func GetLyrics(ctx context.Context, client disgolink.RestClient, track string, skipTrackSource bool) (*Lyrics, error) {
	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/v4/lyrics"), nil)
	if err != nil {
		return nil, err
	}

	q := rq.URL.Query()
	q.Add("track", track)
	q.Add("skipTrackSource", fmt.Sprint(skipTrackSource))
	rq.URL.RawQuery = q.Encode()

	rs, err := client.Do(rq)
	if err != nil {
		return nil, err
	}
	defer rs.Body.Close()

	if rs.StatusCode < 200 || rs.StatusCode >= 300 {
		var lavalinkError lavalink.Error
		if err = json.NewDecoder(rs.Body).Decode(&lavalinkError); err != nil {
			return nil, err
		}
		return nil, lavalinkError
	}

	if rs.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	var l Lyrics
	if err = json.NewDecoder(rs.Body).Decode(&l); err != nil {
		return nil, err
	}
	return &l, nil
}
