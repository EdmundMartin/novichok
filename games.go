package novichok

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const gameArchiveUrl = "https://api.chess.com/pub/player/%s/games/%d/%s"

type PlayerDetails struct {
	Username string `json:"username"`
	Rating   int    `json:"rating"`
	Result   string `json:"result"`
	Id       string `json:"id"`
}

type Game struct {
	Url         string        `json:"url"`
	FeN         string        `json:"fen"`
	Pgn         string        `json:"pgn"`
	StartTime   int           `json:"start_time"`
	EndTime     int           `json:"end_time"`
	TimeControl string        `json:"time_control"`
	Rules       string        `json:"rules"`
	Eco         string        `json:"eco"`
	Tournament  string        `json:"tournament"`
	Match       string        `json:"match"`
	White       PlayerDetails `json:"white"`
	Black       PlayerDetails `json:"black"`
}

type GameArchive struct {
	Games []Game `json:"games"`
}

func (c *ChessComClient) GetGameArchive(ctx context.Context, screenName string, archiveMonth time.Time) (*GameArchive, error) {
	req := &http.Request{URL: urlFromString(fmt.Sprintf(gameArchiveUrl, screenName, archiveMonth.Year(), stringMonthFromTime(archiveMonth)))}
	resp, _, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var archive GameArchive
	err = json.NewDecoder(resp.Body).Decode(&archive)
	if err != nil {
		return nil, err
	}
	return &archive, nil
}
