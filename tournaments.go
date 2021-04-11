package novichok

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const tournamentSummaryUrl = "https://api.chess.com/pub/player/%s/tournaments"

type Tournament struct {
	Url          string `json:"url"`
	Id           string `json:"@id"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Draws        int    `json:"draws"`
	Placement    int    `json:"placement"`
	Status       string `json:"status"`
	TotalPlayers int    `json:"total_players"`
	TimeClass    string `json:"time_class"`
	Type         string `json:"type"`
}

type TournamentSummary struct {
	Finished   []Tournament `json:"finished"`
	InProgress []Tournament `json:"in_progress"`
	Registered []Tournament `json:"registered"`
}

func (t Tournament) String() string {
	return fmt.Sprintf("Tournament{URL: %s, Wins: %d, Losses: %d, Draws: %d}", t.Url, t.Wins, t.Losses, t.Draws)
}

func (c *ChessComClient) GetTournaments(ctx context.Context, screenName string) (*TournamentSummary, error) {
	req := &http.Request{URL: urlFromString(fmt.Sprintf(tournamentSummaryUrl, screenName))}
	resp, _, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var tournaments TournamentSummary
	err = json.NewDecoder(resp.Body).Decode(&tournaments)
	if err != nil {
		return nil, err
	}
	return &tournaments, nil
}
