package services

import (
	// "backend/loaders/db"
	"backend/loaders/db"
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
	db.CardModel.Model(&topics[body.TopicId-1].Cards[body.CardId-1]).Update("opened", true)

	s.topicEvent.SetCurrentCard(topics[body.TopicId-1].Cards[body.CardId-1])

	return topics, nil
}

func (s *topicService) GetCardConn() *extend.ConnModel {
	return s.topicEvent.GetCardConn()
}

func (s *topicService) GetTopics() []*database.Topic {
	return s.topicEvent.GetTopics()
}
