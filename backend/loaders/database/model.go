package database

import (
	"backend/types/database"
	"gorm.io/gorm"
)

var TopicModel *gorm.DB
var TeamModel *gorm.DB
var ScoreModel *gorm.DB
var RawModel *gorm.DB
var QuestionModel *gorm.DB
var CardModel *gorm.DB

func assignModel() {
	TopicModel = DB.Model(new(database.Topic))
	TeamModel = DB.Model(new(database.Team))
	ScoreModel = DB.Model(new(database.Score))
	RawModel = DB.Model(new(database.Raw))
	QuestionModel = DB.Model(new(database.Question))
	CardModel = DB.Model(new(database.Card))

}