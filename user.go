package joomla

import (
	"context"
	"crypto/md5"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
)

var ctx context.Context

func ValidUser(uname string) bool {

	qs := fmt.Sprintf("SELECT username,block FROM %susers WHERE username = ?", Prefix())

	stmt, err := DB.Prepare(qs)
	defer stmt.Close()

	if err != nil {
		return false
	}

	var username string
	var block int

	err = stmt.QueryRow(uname).Scan(&username, &block)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		return false
	case block == 1:
		return false
	default:
		return true
	}

}

func ValidCredentials(uname string, pass string) bool {

	qry := PrepareSQL(`SELECT username,password,block FROM #__users WHERE username = ?`)

	stmt, err := DB.Prepare(qry)

	defer stmt.Close()

	if err != nil {
		return false
	}

	var username string
	var password string
	var block int

	err = stmt.QueryRow(uname).Scan(&username, &password, &block)
	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		return false
	}

	if block == 1 {
		return false
	}

	if len(password) == 32 {
		//this is a md5 hash
		h := md5.New()
		io.WriteString(h, pass)
		if password == fmt.Sprintf("%x", h.Sum(nil)) {
			return true
		}

	}

	berr := bcrypt.CompareHashAndPassword([]byte(password), []byte(pass))
	return berr == nil

}

func GetUser(uname string) (*User, error) {

	var u User

	stmt, err := DB.Prepare(fmt.Sprintf("SELECT id,username,name,email,block,sendEmail,registerDate,lastvisitDate,activation,params,lastResetTime,resetCount,otpKey,otep,requireReset FROM %susers WHERE username = ?", Prefix()))
	defer stmt.Close()

	if err != nil {
		return &u, err
	}

	err = stmt.QueryRow(uname).Scan(&u.ID, &u.Username, &u.Name, &u.Email, &u.Block, &u.SendEmail, &u.RegisterDate, &u.LastvisitDate, &u.Activation, &u.Params, &u.LastResetTime, &u.ResetCount, &u.OtpKey, &u.Otep, &u.RequireReset)

	if err != nil {

		return &u, err

	}

	return &u, nil

}
