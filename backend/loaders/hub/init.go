package hub

import (
	"backend/loaders/db"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"

	"backend/types/database"
	"backend/types/enum"
	"backend/types/extend"
	"backend/utils/logger"
)

func Init() {
	Hub = &Model{
		Topics:      nil,
		Teams:       nil,
		Turned:      nil,
		CurrentCard: nil,
		AdminConn: &extend.ConnModel{
			Context: "ADMIN_CONN",
			Conn:    nil,
			Mutex:   &sync.Mutex{},
		},
		LeaderboardProjectorConn: &extend.ConnModel{
			Context: "LEADERBOARD_PROJECTOR_CONN",
			Conn:    nil,
			Mutex:   &sync.Mutex{},
		},
		CardProjectorConn: &extend.ConnModel{
			Context: "CARD_PROJECTOR_CONN",
			Conn:    nil,
			Mutex:   &sync.Mutex{},
		},
		StudentConns: nil,
		Mode:         enum.ModePreview,
		PreviewCount: 0,
		FinalCandidates: nil,
	}
	if err := Load(); err != nil {
		logger.Log(logrus.Panic, "UNABLE TO LOAD HUB: "+err.Error())
	}
	Watch()
}

func Load() error {
	// * Unmarshal file
	var teams []*database.Team
	var topic []*database.Topic
	var turn []*database.Turn

	if err := db.TopicModel.Preload("Cards").Find(&topic).Error; err != nil {
		return err
	}

	if err := db.TeamModel.Preload("Scores").Find(&teams).Error; err != nil {
		return err
	}
	if err := db.TurnedModel.Find(&turn).Error; err != nil {
		return err
	}
	// * Assign hub
	Hub.Topics = topic
	Hub.Teams = teams


	if len(turn) == 0 {
		Hub.Turned = []*database.Team{
			Hub.Teams[0],
		}
		firstturn := database.Turn{TeamId: &Hub.Teams[0].Id}
		result := db.DB.Create(&firstturn)
		if result.Error != nil {
			return result.Error
		}
	}
	if len(turn) != 0 {
		var teamturn []*uint64
		for _, turned := range turn {
			teamturn = append(teamturn, turned.TeamId)
		}
		if err := db.DB.Find(&teams, teamturn).Error; err != nil {
			return err
		}
		Hub.Turned = teams
	}
	println(len(Hub.Turned))
	logger.Log(logrus.Debug, "LOADED HUB DATA")

	return nil
}

func Watch() {
	// * Create new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Log(logrus.Panic, "UNABLE TO START WATCHER: "+err.Error())
	}

	// * Start listening for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {
					if err := Load(); err != nil {
						logger.Log(logrus.Warn, "UNABLE TO LOAD HUB: "+err.Error())
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Log(logrus.Warn, "WATCHER ERROR: "+err.Error())
			}
		}
	}()

	// * Add a watch path
	err = watcher.Add("./data-reload.json")
	if err != nil {
		log.Fatal(err)
	}
}

func Snapshot() {
	file, err := json.MarshalIndent(Hub, "", "  ")
	if err != nil {
		logger.Log(logrus.Warn, "UNABLE TO MARSHAL SNAPSHOT HUB: "+err.Error())
	}

	if err := os.WriteFile("./data-snap.json", file, 0644); err != nil {
		logger.Log(logrus.Warn, "UNABLE TO WRITE SNAPSHOT HUB: "+err.Error())
	}
}
