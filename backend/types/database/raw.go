package database

type Raw struct {
	TopicId uint64 `json:"topics" gorm:"primaryKey;foreignKey:TopicId;references:Id"`
	Topic   *Topic `json:"topic" gorm:"not null"`
	TeamId  uint64 `json:"teams" gorm:"primaryKey;foreignKey:TeamId;references:Id"`
	Team    *Team  `json:"team" gorm:"not null"`
}
