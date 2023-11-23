package services

import (
	// "backend/loaders/db"
	"backend/loaders/db"
	"backend/loaders/hub"
	"backend/repository"
	"backend/types/database"
	"backend/types/extend"
	"backend/types/payload"
	"backend/types/response"
)

type topicService struct {
	topicEvent repository.TopicRepository
}

func NewTopicService(topicRepository repository.TopicRepository) *topicService {
	return &topicService{topicEvent: topicRepository}
}
func (s *topicService) OpenCard(body *payload.OpenCard) ([]*database.Topic, error) {
	hub.Skip <- false
	topics := s.topicEvent.GetTopics()

	//handle Topic not found
	if len(topics) < int(body.TopicId) {
		return nil, &response.Error{
			Message: "Topic not found",
		}
	}
	if s.topicEvent.GetCurrentCard() != nil {
		return nil, &response.Error{
			Message: "Opened card remaining",
		}
	}

	if topics[body.TopicId-1].Cards[body.CardId-1].Opened {
		return nil, &response.Error{
			Message: "The card has already opened",
		}
	}
	//update card on db
	//chang 3 to 4 if use 4 cards

	cardId := (body.TopicId * 4) - (4 - body.CardId)
	card := &database.Card{Id: &cardId}
	err := db.DB.Model(card).Update("opened", true).Error
	if err != nil {
		println("Failed to update card:", err)
	}
	s.topicEvent.SetCurrentCard(topics[body.TopicId-1].Cards[body.CardId-1])

	return topics, nil
}

func (s *topicService) GetCardConn() *extend.ConnModel {
	return s.topicEvent.GetCardConn()
}

func (s *topicService) GetTopics() []*database.Topic {
	return s.topicEvent.GetTopics()
}
