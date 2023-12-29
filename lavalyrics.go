package lavalyrics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

type Lyrics struct {
	SourceName string           `json:"sourceName"`
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

func GetLyrics(ctx context.Context, client disgolink.RestClient, sessionID string, guildID snowflake.ID) (*Lyrics, error) {
	rq, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/v4/sessions/%s/players/%s/track/lyrics", sessionID, guildID), nil)
	if err != nil {
		return nil, err
	}

	rq.Header.Add("Content-Type", "application/json")

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
