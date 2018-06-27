package models

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/utils"
)

// db.sensores.aggregate([{$unwind:"$tipo_sensor"},{$project:{_id:0,tipo_sensor:1,"prefijo":1}}])
// db.users.find({},{_id:0,id_user:1}).pretty()

type Login struct {
	Email    string `json:"email"`
	Password string `json:"pwd"`
	Source   string `json:"src"`
}

func init() {}

/* curl -d "id_usuario=U0004&nombre=Frank&apellido=Moreno&documento=72972154&username=can&password=123456&f_registro=2017-01-31&h_registro=07:11:00" http://localhost:6568/beagons/users
 */
func VerifyLogin(user Login) (result utils.Response, err error) {
	response := utils.Response{}

	_sn := getDBSession().Copy()
	defer _sn.Close()
	var sessionid = 0
	c := col(_sn, db.ColUsers)

	var exists datamodel.UserSelect

	var query = bson.M{}
	if utils.ValidateEmail(user.Email) {
		query = bson.M{"email": user.Email}

	} else {
		query = bson.M{"mob": user.Email}
	}

	err = c.Find(query).One(&exists)

	if err != nil && err.Error() == "not found" {
		response.Message = "Invalid Login"

		response.Status = false
		response.StausCode = 2
	} else if !exists.Active {
		response.Message = "OTP Not verified!"
		response.Data = exists.Mobile
		response.Status = false
		response.StausCode = 3
	} else {
		//update session in main table
		sessionid = GetNextSequence(_sn, "session")
		response.Message = "Login Successfully"
		response.Status = true
		response.StausCode = 1
		exists.Session = sessionid
		response.Data = exists
	}

	//insert log entry for login
	session := datamodel.UserLoginLog{Email: user.Email, UID: exists.ID, LoginTime: time.Now(),
		ID: sessionid, Source: user.Source, Extra: response}
	logc := col(_sn, db.ColLoginLog)
	_ = logc.Insert(session)

	return response, nil
}

func Logout(user Session) (result utils.Response, err error) {
	response := utils.Response{}

	_sn := getDBSession().Copy()
	defer _sn.Close()

	logc := col(_sn, db.ColLoginLog)

	var exists datamodel.UserLoginLog
	logc.FindId(user.Sessionid).One(&exists)

	if exists.ID > 0 && exists.LogoutTime.Year() > 0001 {
		fmt.Println(exists.LogoutTime.Year())
		response.Message = "Logout Failed. Already logged out."
		response.Status = false
		response.StausCode = 2
		return response, nil
	}

	err = logc.UpdateId(user.Sessionid, bson.M{"logouttime": time.Now()})

	if err != nil && err.Error() == "not found" {
		response.Message = "Logout Failed. Invalid sessionid"
		response.Status = false
		response.StausCode = 0
	} else {
		response.Message = "Logout Successfully"
		response.Status = true
		response.StausCode = 1
	}

	return response, nil
}

type Session struct {
	Sessionid int    `json:"sessionid"`
	Email     string `json:"email"`
	Source    string `json:"src"`
}

func VerifySession(user Session) (result utils.Response, err error) {
	response := utils.Response{}

	_sn := getDBSession().Copy()
	defer _sn.Close()

	logc := col(_sn, db.ColLoginLog)
	var exists datamodel.UserLoginLog
	err = logc.Find(bson.M{"_id": user.Sessionid}).One(&exists)

	if err != nil && err.Error() == "not found" || exists.LogoutTime.Year() > 0001 {
		response.Message = "Invalid Login"
		response.Status = false
		response.StausCode = 0
	} else {
		response.Message = "Login Successfully"
		response.Status = true
		response.StausCode = 0
		response.Data = exists.Extra
	}

	return response, nil
}
