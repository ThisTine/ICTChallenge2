package database

import "time"

type Topic struct {
	Id    *uint64 `json:"id" gorm:"primaryKey;not null"`
	Title *string `json:"title" gorm:"not null"`
	Cards []*Card `json:"cards" gorm:"foreignKey:TopicId;references:Id"`
}

type Question struct {
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
}

type Card struct {
	Id       *uint64       `json:"id" gorm:"primaryKey;not null"`
	Score    *int32        `json:"score" gorm:"not null"`
	Opened   bool          `json:"opened" gorm:"not null"`
	Bonus    bool          `json:"bonus" gorm:"not null"`
	Title    *string       `json:"title" gorm:"not null"`
	TopicId  *uint64       `json:"topic_id" gorm:"foreignKey:TopicId;references:Id;not null"`
	Topic    *Topic        `json:"topic" gorm:"not null"`
	ImageUrl *string       `json:"image_url"  gorm:"not null"`
	Duration time.Duration `json:"duration"  gorm:"not null"`
}
