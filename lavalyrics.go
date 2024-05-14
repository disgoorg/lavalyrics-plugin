package lavalyrics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/disgoorg/disgolink/v4/disgolink"
	"github.com/disgoorg/disgolink/v4/lavalink"
	"github.com/disgoorg/json/v2"
)

var (
	_ disgolink.EventPlugins = (*Plugin)(nil)
	_ disgolink.Plugin       = (*Plugin)(nil)
)

func New() *Plugin {
	return NewWithLogger(slog.Default())
}

func NewWithLogger(logger *slog.Logger) *Plugin {
	return &Plugin{
		eventPlugins: []disgolink.EventPlugin{
			&lyricsFoundHandler{
				logger: logger,
			},
			&lyricsNotFoundHandler{
				logger: logger,
			},
			&lyricsLineHandler{
				logger: logger,
			},
		},
	}
}

type Plugin struct {
	eventPlugins []disgolink.EventPlugin
}

func (p *Plugin) EventPlugins() []disgolink.EventPlugin {
	return p.eventPlugins
}

func (p *Plugin) Name() string {
	return "lavalyrics"
}

func (p *Plugin) Version() string {
	return "1.0.0"
}

var _ disgolink.EventPlugin = (*lyricsFoundHandler)(nil)

type lyricsFoundHandler struct {
	logger *slog.Logger
}

func (h *lyricsFoundHandler) Event() lavalink.EventType {
	return EventTypeLyricsFound
}
func (h *lyricsFoundHandler) OnEventInvocation(player disgolink.Player, data []byte) {
	var e LyricsFoundEvent
	if err := json.Unmarshal(data, &e); err != nil {
		h.logger.Error("Failed to unmarshal LyricsFoundEvent", slog.Any("err", err))
		return
	}

	player.Lavalink().EmitEvent(player, e)
}

var _ disgolink.EventPlugin = (*lyricsNotFoundHandler)(nil)

type lyricsNotFoundHandler struct {
	logger *slog.Logger
}

func (h *lyricsNotFoundHandler) Event() lavalink.EventType {
	return EventTypeLyricsNotFound
}

func (h *lyricsNotFoundHandler) OnEventInvocation(player disgolink.Player, data []byte) {
	var e LyricsNotFoundEvent
	if err := json.Unmarshal(data, &e); err != nil {
		h.logger.Error("Failed to unmarshal LyricsNotFoundEvent", err)
		return
	}

	player.Lavalink().EmitEvent(player, e)
}

type lyricsLineHandler struct {
	logger *slog.Logger
}

func (h *lyricsLineHandler) Event() lavalink.EventType {
	return EventTypeLyricsLine
}

func (h *lyricsLineHandler) OnEventInvocation(player disgolink.Player, data []byte) {
	var e LyricsLineEvent
	if err := json.Unmarshal(data, &e); err != nil {
		h.logger.Error("Failed to unmarshal LyricsLineEvent", err)
		return
	}

	player.Lavalink().EmitEvent(player, e)
}
