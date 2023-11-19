package database

type Team struct {
	Id     uint64   `json:"id" gorm:"primaryKey;not null"`
	Name   string   `json:"name" gorm:"primaryKey;not null"`
	School string   `json:"school" gorm:"primaryKey;not null"`
	Token  string   `json:"token,omitempty"`
	Scores []*Score `json:"scores" gorm:"foreignKey:TeamId;references:Id"`
}
