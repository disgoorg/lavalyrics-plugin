package lavalyrics

import (
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

const (
	EventTypeLyricsFound    lavalink.EventType = "LyricsFoundEvent"
	EventTypeLyricsNotFound lavalink.EventType = "LyricsNotFoundEvent"
	EventTypeLyricsLine     lavalink.EventType = "LyricsLineEvent"
)

type LavaLyricsEventListener interface {
	OnLyricsFound(player disgolink.Player, event LyricsFoundEvent)
	OnLyricsNotFound(player disgolink.Player, event LyricsNotFoundEvent)
	OnLyricsLine(player disgolink.Player, event LyricsLineEvent)
}

type LyricsFoundEvent struct {
	GuildID_ snowflake.ID `json:"guildId"`
	Lyrics   Lyrics       `json:"lyrics"`
}

func (LyricsFoundEvent) Op() lavalink.Op {
	return lavalink.OpEvent
}

func (LyricsFoundEvent) Type() lavalink.EventType {
	return EventTypeLyricsFound
}

func (e LyricsFoundEvent) GuildID() snowflake.ID {
	return e.GuildID_
}

type LyricsNotFoundEvent struct {
	GuildID_ snowflake.ID `json:"guildId"`
}

func (LyricsNotFoundEvent) Op() lavalink.Op {
	return lavalink.OpEvent
}

func (LyricsNotFoundEvent) Type() lavalink.EventType {
	return EventTypeLyricsNotFound
}

func (e LyricsNotFoundEvent) GuildID() snowflake.ID {
	return e.GuildID_
}

type LyricsLineEvent struct {
	GuildID_  snowflake.ID `json:"guildId"`
	LineIndex int          `json:"lineIndex"`
	Line      Line         `json:"line"`
	Skipped   bool         `json:"skipped"`
}

func (LyricsLineEvent) Op() lavalink.Op {
	return lavalink.OpEvent
}

func (LyricsLineEvent) Type() lavalink.EventType {
	return EventTypeLyricsLine
}

func (e LyricsLineEvent) GuildID() snowflake.ID {
	return e.GuildID_
}
