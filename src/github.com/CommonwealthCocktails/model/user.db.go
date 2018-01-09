// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/user.db.go:package model
package model

import (
	"bytes"
	"github.com/CommonwealthCocktails/connectors"
	"github.com/golang/glog"
	"html"
	"strconv"
	"strings"
	"time"
)

//CREATE, UPDATE, DELETE
//Insert a user record into the database
func (user *User) InsertUser() int {
	//set the ID to zero to indicate an insert
	user.ID = 0
	return user.processUser()
}

//Update a user record in the database based on ID
func (user *User) UpdateUser() int {
	return user.processUser()
}

//process a user record into the database for update and insert
func (user *User) processUser() int {
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

	//If the ID is zero then do an insert else do an update based on the ID
	if user.ID == 0 {
		buffer.WriteString("INSERT INTO `users` ( ")
	} else {
		buffer.WriteString("UPDATE `users` SET ")
	}

	//Append the correct columns to be added based on data available in the
	//data structure
	if user.Username != "" {
		if user.ID == 0 {
			buffer.WriteString("`userName`,")
		} else {
			buffer.WriteString("`userName`=?,")
		}
		args = append(args, html.EscapeString(user.Username))
	}
	if user.Password != "" {
		if user.ID == 0 {
			buffer.WriteString("`userPassword`,")
		} else {
			buffer.WriteString("`userPassword`=?,")
		}
		args = append(args, user.Password)
	}
	if user.Email != "" {
		if user.ID == 0 {
			buffer.WriteString("`userEmail`,")
		} else {
			buffer.WriteString("`userEmail`=?,")
		}
		args = append(args, user.Email)
	}
	if user.VerificationCode != "" {
		if user.ID == 0 {
			buffer.WriteString("`userVerificationCode`,")
		} else {
			buffer.WriteString("`userVerificationCode`=?,")
		}
		args = append(args, user.VerificationCode)
	}
	if user.FirstName != "" {
		if user.ID == 0 {
			buffer.WriteString("`userFirstName`,")
		} else {
			buffer.WriteString("`userFirstName`=?,")
		}
		args = append(args, user.FirstName)
	}
	if user.LastName != "" {
		if user.ID == 0 {
			buffer.WriteString("`userLastName`,")
		} else {
			buffer.WriteString("`userLastName`=?,")
		}
		args = append(args, user.LastName)
	}
	if user.ID == 0 {
		buffer.WriteString("`userVerificationInitTime`,")
	} else {
		buffer.WriteString("`userVerificationInitTime`=?,")
	}
	args = append(args, user.VerificationInitTime)
	if user.ID == 0 {
		buffer.WriteString("`userVerificationComplete`,")
	} else {
		buffer.WriteString("`userVerificationComplete`=?,")
	}
	args = append(args, user.VerificationComplete)

	//Cleanup the query and append where if it is an update
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	if user.ID == 0 {
		vals := strings.Repeat("?,", len(args))
		vals = strings.TrimRight(vals, ",")
		query = query + ") VALUES (" + vals + ");"
	} else {
		query = query + " WHERE `idUser`=?;"
		args = append(args, user.ID)
	}

	//Lets do this thing
	glog.Infoln(query)
	ret := -1
	r, err := conn.Exec(query, args...)
	if err != nil {
		glog.Error(err)
		return ret
	}
	id, _ := r.LastInsertId()
	ret = int(id)
	return ret
}

//Insert a user session record into the database
func (userSession *UserSession) InsertUserSession() int {
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

	buffer.WriteString("INSERT INTO `usersessions` ( ")

	//Append the correct columns to be added based on data available in the
	//data structure
	if userSession.SessionKey != "" {
		buffer.WriteString("`usersessionSessionKey`,")
		args = append(args, userSession.SessionKey)
	} else {
		return -1
	}
	buffer.WriteString("`idUser`,")
	args = append(args, strconv.Itoa(userSession.User.ID))
	buffer.WriteString("`usersessionCSRFBase`,")
	args = append(args, userSession.CSRFBase)
	buffer.WriteString("`usersessionCSRFKey`,")
	glog.Infoln(userSession.CSRFKey)
	args = append(args, userSession.CSRFKey)
	buffer.WriteString("`usersessionCSRFGenTime`,")
	args = append(args, userSession.CSRFGenTime.Format(time.RFC3339))
	buffer.WriteString("`usersessionLoginTime`,")
	args = append(args, userSession.LoginTime.Format(time.RFC3339))
	buffer.WriteString("`usersessionLastSeenTime`,")
	args = append(args, userSession.LastSeenTime.Format(time.RFC3339))
	buffer.WriteString("`usersessionIsDefaultUser`,")
	args = append(args, btoi(userSession.IsDefaultUser))

	//Cleanup the query and append where if it is an update
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	vals := strings.Repeat("?,", len(args))
	vals = strings.TrimRight(vals, ",")
	query = query + ") VALUES (" + vals + ");"

	//Lets do this thing
	glog.Infoln(query)
	ret := -1
	r, err := conn.Exec(query, args...)
	if err != nil {
		glog.Error(err)
		return ret
	}
	id, _ := r.LastInsertId()
	ret = int(id)
	return ret
}

//Update a user session record in the database based on previous session id if
//not empty else based on the user session object's session key.
func (userSession *UserSession) UpdateUserSession(prevSessionKey string) int {
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

	buffer.WriteString("UPDATE `usersessions` SET ")

	//Append the correct columns to be added based on data available in the
	//data structure
	if userSession.SessionKey != "" {
		buffer.WriteString("`usersessionSessionKey`=?,")
		args = append(args, userSession.SessionKey)
	} else {
		return -1
	}
	buffer.WriteString("`idUser`=?,")
	args = append(args, strconv.Itoa(userSession.User.ID))
	buffer.WriteString("`usersessionCSRFBase`=?,")
	args = append(args, userSession.CSRFBase)
	buffer.WriteString("`usersessionCSRFKey`=?,")
	args = append(args, userSession.CSRFKey)
	buffer.WriteString("`usersessionCSRFGenTime`=?,")
	args = append(args, userSession.CSRFGenTime.Format(time.RFC3339))
	buffer.WriteString("`usersessionLoginTime`=?,")
	args = append(args, userSession.LoginTime.Format(time.RFC3339))
	buffer.WriteString("`usersessionLastSeenTime`=?,")
	args = append(args, userSession.LastSeenTime.Format(time.RFC3339))
	buffer.WriteString("`usersessionIsDefaultUser`=?,")
	args = append(args, btoi(userSession.IsDefaultUser))

	//Cleanup the query and append where if it is an update
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	query = query + " WHERE `usersessionSessionKey`=?;"
	if prevSessionKey != "" {
		args = append(args, prevSessionKey)
	} else {
		args = append(args, userSession.SessionKey)
	}

	//Lets do this thing
	glog.Infoln(query)
	ret := -1
	r, err := conn.Exec(query, args...)
	if err != nil {
		glog.Error(err)
		return ret
	}
	id, _ := r.LastInsertId()
	ret = int(id)
	return ret
}

//Put an OAuth Key and time created into the database
func (oauth *OAuth) InsertOAuth() int {
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

	buffer.WriteString("INSERT INTO `oauth` ( ")

	//Append the correct columns to be added based on data available in the
	//data structure
	if oauth.Key != "" {
		buffer.WriteString("`oauthKey`,")
		args = append(args, oauth.Key)
	} else {
		return -1
	}
	buffer.WriteString("`oauthTime`,")
	args = append(args, oauth.Time.Format(time.RFC3339))

	//Cleanup the query and append where if it is an update
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	vals := strings.Repeat("?,", len(args))
	vals = strings.TrimRight(vals, ",")
	query = query + ") VALUES (" + vals + ");"

	//Lets do this thing
	glog.Infoln(query)
	ret := -1
	r, err := conn.Exec(query, args...)
	if err != nil {
		glog.Error(err)
		return ret
	}
	id, _ := r.LastInsertId()
	ret = int(id)
	return ret
}

//Remove an OAuth Key and time from the database by OAuth Key
func (oauth *OAuth) DeleteOAuth() int {
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

	buffer.WriteString("DELETE FROM `oauth` ")

	buffer.WriteString("WHERE `oauthKey`=?")
	args = append(args, oauth.Key)

	//Cleanup the query and append where if it is an update
	query := buffer.String()
	query = strings.TrimRight(query, ",")
	query = query + ";"

	//Lets do this thing
	glog.Infoln(query)
	ret := -1
	r, err := conn.Exec(query, args...)
	if err != nil {
		glog.Error(err)
		return ret
	}
	aff, _ := r.RowsAffected()
	ret = int(aff)
	return ret
}

//SELECTS
//Get a user from the user id and the user name both of which should be
//unique.
func (user *User) SelectUser() *User {
	var ret User
	conn, _ := connectors.GetDB()
	glog.Infoln(user.Username)
	var buffer bytes.Buffer
	var canQuery = false
	var args []interface{}
	buffer.WriteString("SELECT `idUser`, `userName`, `userPassword`, `userEmail`, `userFirstName`, `userLastName`, `userVerificationInitTime`, `userVerificationCode`, `userVerificationComplete`, `userRole` FROM `users` WHERE ")
	if user.ID != 0 {
		buffer.WriteString("`idUser`=? AND ")
		args = append(args, strconv.Itoa(user.ID))
		canQuery = true
	}
	if user.Username != "" {
		buffer.WriteString("`userName`=? AND ")
		args = append(args, html.EscapeString(user.Username))
		canQuery = true
	}

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&ret.ID, &ret.Username, &ret.Password, &ret.Email, &ret.FirstName, &ret.LastName, &ret.VerificationInitTime, &ret.VerificationCode, &ret.VerificationComplete, &ret.Role)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln(ret.ID, ret.Username, ret.Password, ret.Email, ret.FirstName, ret.LastName, ret.VerificationInitTime, ret.VerificationCode, ret.VerificationComplete, ret.Role)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return &ret
}

//Get a user from the the user name or the email
func (user *User) SelectUsersByIdORUsernameOREmail() []User {
	var ret []User
	conn, _ := connectors.GetDB()
	glog.Infoln(user.Username)
	var buffer bytes.Buffer
	var canQuery = false
	var args []interface{}
	buffer.WriteString("SELECT `idUser`, `userName`, `userPassword`, `userEmail`, `userFirstName`, `userLastName`, `userVerificationInitTime`, `userVerificationCode`, `userVerificationComplete`, `userRole` FROM `users` WHERE ")
	if user.ID != 0 {
		buffer.WriteString("`idUser`=? OR ")
		args = append(args, strconv.Itoa(user.ID))
		canQuery = true
	}
	if user.Username != "" {
		buffer.WriteString("`userName`=? OR ")
		args = append(args, html.EscapeString(user.Username))
		canQuery = true
	}
	if user.Email != "" {
		buffer.WriteString("`userEmail`=? OR ")
		args = append(args, html.EscapeString(user.Email))
		canQuery = true
	}

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " OR")
		query = query + ";"
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
		}
		defer rows.Close()
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.VerificationInitTime, &user.VerificationCode, &user.VerificationComplete, &user.Role)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln(user.ID, user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.VerificationInitTime, user.VerificationCode, user.VerificationComplete, user.Role)
			ret = append(ret, user)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return ret
}

//Check the database for a user based on verification code and email address
func (user *User) SelectUserForVerification() *User {
	var ret User
	conn, _ := connectors.GetDB()
	glog.Infoln(user.Username)
	var buffer bytes.Buffer
	var canQuery = false
	var args []interface{}
	buffer.WriteString("SELECT `idUser`, `userName`, `userPassword`, `userEmail`, `userFirstName`, `userLastName`, `userVerificationInitTime`, `userVerificationCode`, `userVerificationComplete`, `userRole` FROM `users` WHERE ")
	if user.VerificationCode != "" {
		buffer.WriteString(" `userVerificationCode`=? AND")
		args = append(args, html.EscapeString(user.VerificationCode))
		canQuery = true
	}
	if user.Email != "" {
		buffer.WriteString("`userEmail`=? AND")
		args = append(args, user.Email)
		canQuery = true
	}
	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
			return nil
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&ret.ID, &ret.Username, &ret.Password, &ret.Email, &ret.FirstName, &ret.LastName, &ret.VerificationInitTime, &ret.VerificationCode, &ret.VerificationComplete, &ret.Role)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln(ret.ID, ret.Username, ret.Password, ret.Email, ret.FirstName, ret.LastName, ret.VerificationInitTime, ret.VerificationCode, ret.VerificationComplete, ret.Role)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return &ret
}

//Check the database for a user based on user name and password, if Ouath is
//set then check the user based on email
func (user *User) SelectUserForLogin(isOauth bool) *User {
	var ret User
	conn, _ := connectors.GetDB()
	glog.Infoln(user.Username)
	var buffer bytes.Buffer
	var canQuery = false
	var args []interface{}
	buffer.WriteString("SELECT `idUser`, `userName`, `userPassword`, `userEmail`, `userFirstName`, `userLastName`, `userVerificationInitTime`, `userVerificationCode`, `userVerificationComplete`, `userRole` FROM `users` WHERE ")
	if user.Username != "" {
		buffer.WriteString(" `userName`=? AND")
		args = append(args, html.EscapeString(user.Username))
		canQuery = true
	}
	if isOauth {
		if user.Email != "" {
			buffer.WriteString("`userEmail`=? AND")
			args = append(args, user.Email)
			canQuery = true
		}
	}

	if canQuery {
		query := buffer.String()
		query = strings.TrimRight(query, " AND")
		query = query + ";"
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
			return nil
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&ret.ID, &ret.Username, &ret.Password, &ret.Email, &ret.FirstName, &ret.LastName, &ret.VerificationInitTime, &ret.VerificationCode, &ret.VerificationComplete, &ret.Role)
			if err != nil {
				glog.Error(err)
			}
			glog.Infoln(ret.ID, ret.Username, ret.Password, ret.Email, ret.FirstName, ret.LastName, ret.VerificationInitTime, ret.VerificationCode, ret.VerificationComplete, ret.Role)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
		}
	}
	return &ret
}

//Get a user session based on the session key and only the session key.  Not
//this handles the default user case also.
//TODO: right now this calls the select user function.  This should really get
//the user information all in one query so I don't have to hit the DB twice
//each time a user goes to a new page.
func (us *UserSession) SelectUserSession() []UserSession {
	var ret []UserSession
	conn, _ := connectors.GetDB()
	var args []interface{} //arguments for variables in the data struct
	var buffer bytes.Buffer
	var canQuery = false
	buffer.WriteString("SELECT `idUser`, `usersessionSessionKey`, `usersessionCSRFGenTime`, `usersessionCSRFBase`, `usersessionCSRFKey`, `usersessionLoginTime`, `usersessionLastSeenTime`, `usersessionIsDefaultUser` FROM `usersessions` WHERE ")
	if us.SessionKey != "" {
		buffer.WriteString("`usersessionSessionKey`=?;")
		args = append(args, us.SessionKey)
		canQuery = true
	}
	if canQuery {
		query := buffer.String()
		glog.Infoln(query)
		rows, err := conn.Query(query, args...)
		if err != nil {
			glog.Error(err)
			return ret
		}
		defer rows.Close()
		for rows.Next() {
			var us UserSession
			var isDefaultUser string
			err := rows.Scan(&us.User.ID, &us.SessionKey, &us.CSRFGenTime, &us.CSRFBase, &us.CSRFKey, &us.LoginTime, &us.LastSeenTime, &isDefaultUser)
			if err != nil {
				glog.Error(err)
			}
			us.IsDefaultUser, _ = strconv.ParseBool(isDefaultUser)
			if us.User.ID != 0 && !us.IsDefaultUser {
				us.User = *us.User.SelectUser()
			}
			ret = append(ret, us)
			glog.Infoln(us.User.ID, us.SessionKey, us.CSRFBase, us.CSRFKey, us.CSRFGenTime, us.LoginTime, us.LastSeenTime, us.IsDefaultUser)
		}
		err = rows.Err()
		if err != nil {
			glog.Error(err)
			return ret
		}
	}
	return ret
}

//Get a OAuth key out of the database with the time the key was created
func (oauth *OAuth) SelectOAuthByKey() *OAuth {
	var ret OAuth
	conn, _ := connectors.GetDB() //get db connection
	var args []interface{}        //arguments for variables in the data struct
	var buffer bytes.Buffer       //buffer for the query

	buffer.WriteString("SELECT `oauthKey`, `oauthTime` FROM `oauth` ")

	if oauth.Key != "" {
		buffer.WriteString("WHERE `oauthKey`=?")
		args = append(args, oauth.Key)
	} else {
		return nil
	}

	query := buffer.String()
	query = query + ";"
	glog.Infoln(query)
	rows, err := conn.Query(query, args...)
	if err != nil {
		glog.Error(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ret.Key, &ret.Time)
		if err != nil {
			glog.Error(err)
		}
		glog.Infoln(ret.Key, ret.Time)
	}
	err = rows.Err()
	if err != nil {
		glog.Error(err)
	}

	return &ret
}
