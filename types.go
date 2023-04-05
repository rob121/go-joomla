package joomla

import (
	"fmt"
	"time"
)

type User struct {
	ID            int
	Name          string
	Username      string
	Email         string
	Password      string
	Block         int
	SendEmail     int
	RegisterDate  time.Time
	LastvisitDate time.Time
	Activation    string
	Params        string
	LastResetTime time.Time
	ResetCount    int
	OtpKey        string
	Otep          string
	RequireReset  int
}

func (u *User) Groups() ([]UserGroup, error) {

	var group []UserGroup

	// Execute the query
	results, err := DB.Query(fmt.Sprintf(`SELECT user_id,group_id,title FROM %suser_usergroup_map  INNER JOIN %susergroups g ON (group_id = g.id) WHERE user_id=?`, Prefix(), Prefix()), u.ID)

	if err != nil {
		return group, err
	}

	for results.Next() {
		var ug UserGroup
		// for each row, scan the result into our tag composite object
		err = results.Scan(&ug.UserID, &ug.GroupID, &ug.Title)
		if err != nil {
			return group, nil
		}

		group = append(group, ug)

	}

	return group, nil

}

type UserGroup struct {
	UserID  int
	GroupID int
	Title   string
}
