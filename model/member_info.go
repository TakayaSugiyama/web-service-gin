package model

import (
	"database/sql"
	"time"

	"github.com/TakayaSugiyama/web-service-gin/db"
)

type MemberInfo struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	NickName      string    `json:"nickname"`
	ImageLinkURL  string    `json:"image_link_url"`
	BirthdayYear  int       `json:"birthday_year"`
	BirthdayMonth int       `json:"birthday_month"`
	BirthdayDay   int       `json:"birthday_day"`
	BloodType     string    `json:"blood_type"`
	PlaceOfBirth  string    `json:"place_of_birth"`
	Height        int       `json:"height"`
	Hobby         string    `json:"hobby"`
	SpecialSkill  string    `json:"special_skill"`
	BestFeature   string    `json:"best_feature"`
	MemberID      int       `json:"member_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func GetMemberInfos() ([]MemberInfo, error) {
	db := db.ConnectDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM member_infos")
	if err != nil {
		return nil, err
	}
	var members []MemberInfo

	for rows.Next() {
		var member MemberInfo

		if err := rows.Scan(&member.ID, &member.Name, &member.NickName, &member.ImageLinkURL, &member.BirthdayYear, &member.BirthdayMonth, &member.BirthdayDay,
			&member.BloodType, &member.PlaceOfBirth, &member.Height, &member.Hobby, &member.SpecialSkill, &member.BestFeature, &member.MemberID, &member.CreatedAt, &member.UpdatedAt); err != nil {
			panic(err)
		}
		members = append(members, member)
	}
	if err = rows.Err(); err != nil {
		return members, err
	}
	return members, nil
}

func GetMemberInfo(id string) (MemberInfo, error) {
	var member MemberInfo
	db := db.ConnectDB()
	defer db.Close()
	err := db.QueryRow("SELECT * FROM member_infos WHERE id = ?", id).Scan(&member.ID, &member.Name, &member.NickName, &member.ImageLinkURL, &member.BirthdayYear, &member.BirthdayMonth, &member.BirthdayDay,
		&member.BloodType, &member.PlaceOfBirth, &member.Height, &member.Hobby, &member.SpecialSkill, &member.BestFeature, &member.MemberID, &member.CreatedAt, &member.UpdatedAt)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return member, err
}
