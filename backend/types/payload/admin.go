package payload

import "backend/types/enum"

type UpdateScore struct {
	Update [10]int `json:"update"`
}

type LeaderboardMode struct {
	Mode *enum.Mode `json:"mode"`
}
