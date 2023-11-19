package database

type Raw struct {
	TopicId []*Topic `json:"topics" gorm:"primaryKey;foreignKey:TopicId;references:Id"`
	TeamId  []*Team  `json:"teams" gorm:"primaryKey;foreignKey:TeamId;references:Id"`
}
