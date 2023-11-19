package database

type Turned struct {
	TeamId *uint64 `json:"team_id" gorm:"primaryKey;not null"`
}
