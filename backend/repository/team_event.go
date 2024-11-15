package repository

import (
	"backend/loaders/db"
	"backend/loaders/hub"
	"backend/types/database"
	"backend/types/extend"
)

type teamEvent struct {
	hub *hub.Model
}

func NewTeamEvent(hub *hub.Model) *teamEvent {
	return &teamEvent{hub: hub}
}

func (r *teamEvent) GetTeamById(id uint64) *database.Team {
	return nil
}

func (r *teamEvent) GetTeams() []*database.Team {
	var teams []*database.Team
	result := db.TeamModel.Find(&teams)
	if result.Error != nil {
		return nil
	}
	return teams
	// return hub.Hub.Teams
}


func (r *teamEvent) GetTurned() []*database.Team {
	return hub.Hub.Turned
}

func (r *teamEvent) SetTurned(turn []*database.Team) {
	hub.Hub.Turned = turn
}

func (r *teamEvent) GetLeaderboardConn() *extend.ConnModel {
	return hub.Hub.LeaderboardProjectorConn
}

func (r *teamEvent) GetAdminConn() *extend.ConnModel {
	return hub.Hub.AdminConn
}

func (r *teamEvent) GetStudentConns() []*extend.ConnModel {
	return hub.Hub.StudentConns
}

func (r *teamEvent) SetFinalCandidates(candidates []*database.Team) {
	hub.Hub.FinalCandidates = candidates
}
