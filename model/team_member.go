package model

import "github.com/TakayaSugiyama/web-service-gin/db"

type TeamMember struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProfileLink  string `json:"profile_link"`
	TeamName     string `json:"team_name"`
	ImageLinkUrl string `json:"image_link_url"`
}

func GetRandomMembers() ([]TeamMember, error) {
	db := db.ConnectDB()
	defer db.Close()
	var members []TeamMember

	rows, err := db.Query(`
	SELECT
		members.id,
		members.name,
		members.profile_link,
		teams.name,
		member_infos.image_link_url
	FROM
		members
		JOIN teams ON teams.id = members.team_id
		JOIN member_infos ON member_infos.member_id = members.id
	ORDER BY
		RAND()
	LIMIT 10
	`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var member TeamMember

		if err := rows.Scan(&member.ID, &member.Name, &member.ProfileLink, &member.TeamName, &member.ImageLinkUrl); err != nil {
			panic(err)
		}
		members = append(members, member)
	}
	return members, err
}

func GetTeamMembers(id string) ([]TeamMember, error) {
	db := db.ConnectDB()
	defer db.Close()
	var members []TeamMember

	rows, err := db.Query(`
	SELECT
		members.id,
		members.name,
		members.profile_link,
		teams.name,
		member_infos.image_link_url
	FROM
		members
		JOIN teams ON teams.id = members.team_id
		JOIN member_infos ON member_infos.member_id = members.id
	WHERE
		members.team_id = ?
	`, id)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var member TeamMember

		if err := rows.Scan(&member.ID, &member.Name, &member.ProfileLink, &member.TeamName, &member.ImageLinkUrl); err != nil {
			panic(err)
		}
		members = append(members, member)
	}
	return members, err
}
