package model

import (
	"time"
)

type Resp struct {
	State string `json:"state,ommitempty"`
	Clan  Clan   `json:"clan"`
}

type Clan struct {
	Tag          string `json:"tag"`
	Name         string `json:"name"`
	BadgeId      int    `json:"badgeId"`
	Fame         int    `json:"fame"`
	RepairPoints int    `json:"repairPoints"`
	Participants []User `json:"participants"`
	PeriodPoints int    `json:"periodPoints"`
	ClanScore    int    `json:"clanScore"`
}

type User struct {
	Tag            string `json:"tag"`
	Name           string `json:"name"`
	Fame           int    `json:"fame"`
	RepairPoints   int    `json:"repairPoint"`
	BoatAttacks    int    `json:"boatAttacks"`
	DecksUsed      int    `json:"decksUsed"`
	DecksUsedToday int    `json:"decksUsedToday"`
}

type Runaway struct {
	Tag            string `json:"tag"`
	Name           string `json:"name"`
	DecksUsed      int    `json:"decksUsed"`
	DecksUsedToday int    `json:"decksUsedToday"`
}

type Message struct {
	Chat_id int    `json:"chat_id"`
	Text    string `json:"text"`
}

type DbLog struct {
	Year        int
	Week        int
	Day         time.Time
	Player      string
	BattlesDone int
	Comments    string
}

func NewDbLog(year, week int, day time.Time, player string, battlesdone int, comments string) DbLog {
	return DbLog{
		Year:        year,
		Week:        week,
		Day:         day,
		Player:      player,
		BattlesDone: battlesdone,
		Comments:    comments,
	}
}
