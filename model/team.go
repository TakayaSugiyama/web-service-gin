package model

import (
	"fmt"
	"time"

	"github.com/TakayaSugiyama/web-service-gin/db"
)

type Team struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	TeamCount int       `json:"team_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetTeams() ([]Team, error) {
	db := db.ConnectDB()
	rows, err := db.Query("SELECT * FROM teams;")
	if err != nil {
		return nil, err
	}
	var teams []Team

	for rows.Next() {
		var team Team

		if err := rows.Scan(&team.ID, &team.Name, &team.TeamCount, &team.CreatedAt, &team.UpdatedAt); err != nil {
			panic(err)
		}
		teams = append(teams, team)
	}
	if err = rows.Err(); err != nil {
		return teams, err
	}
	fmt.Println(teams)
	return teams, nil
}
