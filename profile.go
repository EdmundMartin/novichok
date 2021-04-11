package novichok

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const playerProfileUrl = "https://api.chess.com/pub/player/%s"
const playerStatsUrl = "https://api.chess.com/pub/player/%s/stats"

type PlayerProfile struct {
	ID         string `json:"@id"`
	URL        string `json:"url"`
	Username   string `json:"username"`
	PlayerID   int64  `json:"player_id"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Location   string `json:"location"`
	Country    string `json:"country"`
	Joined     int    `json:"joined"`
	LastOnline int    `json:"last_online"`
	Followers  int    `json:"followers"`
	IsStreamer bool   `json:"is_streamer"`
	TwitchURL  string `json:"twitch_url"`
	FideRating int    `json:"fide"`
}

func (p PlayerProfile) String() string {
	return fmt.Sprintf("PlayerProfile{Username: %s}", p.Username)
}

type LastResult struct {
	Rating int `json:"rating"`
	Date   int `json:"date"`
	Rd     int `json:"rd"`
}

func (l LastResult) String() string {
	return fmt.Sprintf("LastResult{Rating: %d, Date: %d, Rd: %d}", l.Rating, l.Date, l.Rd)
}

type BestResult struct {
	Rating  int    `json:"rating"`
	Date    int    `json:"date"`
	GameURL string `json:"game"`
}

func (b BestResult) String() string {
	return fmt.Sprintf("BestResult{Rating: %d, Date: %d, GameURL: %s}", b.Rating, b.Date, b.GameURL)
}

type Record struct {
	Won               int `json:"win"`
	Losses            int `json:"loss"`
	Drawn             int `json:"draw"`
	TimePerMove       int `json:"time_per_move"`
	TimeoutPercentage int `json:"timeout_percent"`
}

func (r Record) String() string {
	return fmt.Sprintf("Record{Won: %d, Lost: %d, Drawn: %d}", r.Won, r.Losses, r.Drawn)
}

type GameHistory struct {
	Record Record     `json:"record"`
	Last   LastResult `json:"last"`
	Best   BestResult `json:"best"`
}

type TacticRating struct {
	Rating int `json:"rating"`
	Date   int `json:"date"`
}

func (t TacticRating) String() string {
	return fmt.Sprintf("TacticRating{Rating: %d, Date: %d}", t.Rating, t.Date)
}

type TacticsHistory struct {
	Highest TacticRating `json:"highest"`
	Lowest  TacticRating `json:"lowest"`
}

type PlayerStats struct {
	BlitzStats  GameHistory    `json:"chess_blitz"`
	DailyStats  GameHistory    `json:"chess_daily"`
	RapidStats  GameHistory    `json:"chess_rapid"`
	BulletStats GameHistory    `json:"chess_bullet"`
	TacticStats TacticsHistory `json:"tactics"`
}

func (c *ChessComClient) GetPlayerProfile(ctx context.Context, screenName string) (*PlayerProfile, error) {
	req := &http.Request{URL: urlFromString(fmt.Sprintf(playerProfileUrl, screenName))}
	res, _, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var prof PlayerProfile
	err = json.NewDecoder(res.Body).Decode(&prof)
	if err != nil {
		return nil, err
	}
	return &prof, nil
}

func (c *ChessComClient) GetPlayerStats(ctx context.Context, screenName string) (*PlayerStats, error) {
	req := &http.Request{URL: urlFromString(fmt.Sprintf(playerStatsUrl, screenName))}
	res, _, err := c.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	var stats PlayerStats
	err = json.NewDecoder(res.Body).Decode(&stats)
	if err != nil {
		return nil, err
	}
	return &stats, nil
}
