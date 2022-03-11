package model

import (
	"database/sql"
	"time"

	"github.com/TakayaSugiyama/web-service-gin/db"
)

type Member struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	TeamID      int       `json:"team_id"`
	ProfileLink string    `json:"profile_link"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetMembers() ([]Member, error) {
	db := db.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM members")
	if err != nil {
		return nil, err
	}
	var members []Member

	for rows.Next() {
		var member Member

		if err := rows.Scan(&member.ID, &member.Name, &member.ProfileLink, &member.TeamID, &member.CreatedAt, &member.UpdatedAt); err != nil {
			panic(err)
		}
		members = append(members, member)
	}
	if err = rows.Err(); err != nil {
		return members, err
	}
	return members, nil
}

func GetMember(id string) (Member, error) {
	var member Member
	db := db.ConnectDB()
	defer db.Close()
	err := db.QueryRow("SELECT * FROM members WHERE id = ?", id).Scan(&member.ID, &member.Name, &member.ProfileLink, &member.TeamID, &member.CreatedAt, &member.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return member, err
}
