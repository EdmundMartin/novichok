package novichok

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const titledPlayersUrl = "https://api.chess.com/pub/titled/%s"

type TitledPlayers struct {
	Players []string `json:"players"`
}

var validTitles = map[string]interface{}{
	"GM":  nil,
	"WGM": nil,
	"IM":  nil,
	"WIM": nil,
	"FM":  nil,
	"WFM": nil,
	"NM":  nil,
	"WNM": nil,
	"CM":  nil,
	"WCM": nil,
}

func (c *ChessComClient) GetTitledPlayers(ctx context.Context, title string) (*TitledPlayers, error) {
	if _, ok := validTitles[title]; ok != true {
		return nil, fmt.Errorf("%s is not a valid Chess title", title)
	}
	req := &http.Request{URL: urlFromString(fmt.Sprintf(titledPlayersUrl, title))}
	resp, _, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var titledPlayers TitledPlayers
	err = json.NewDecoder(resp.Body).Decode(&titledPlayers)
	if err != nil {
		return nil, err
	}
	return &titledPlayers, nil
}
