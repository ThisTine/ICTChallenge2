package database

type Score struct {
	Id     *uint64 `json:"id" gorm:"primaryKey;not null"`
	Change *int32  `json:"change" gorm:"not null"`
	Total  *int32  `json:"total" gorm:"not null"`
	TeamId *uint64 `json:"team_id" gorm:"not null"`
	Team   *Team   `json:"team" gorm:"foreignKey:TeamId;references:Id;not null"`
}
