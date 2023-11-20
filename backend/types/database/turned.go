package database

type Turn struct {
	TeamId *uint64 `json:"team_id" gorm:"primaryKey;not null"`
}
