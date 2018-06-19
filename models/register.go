package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/utils"
)

//RegisterUser for registering user
func RegisterUser(user datamodel.User) (response utils.Response, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	/*Pre checking*/

	response = utils.Response{}

	if user.Email == "" {
		response.Message = "Email is required"
		response.Status = false
		response.StausCode = 0
		return response, nil
	}

	user.Created = time.Now()
	user.Updated = time.Now()

	c := col(_sn, db.ColUsers)
	var exist datamodel.User
	err = c.Find(bson.M{"email": user.Email}).One(&exist)
	if err != nil && err.Error() != "not found" {
		response.Message = "Error while user registration. Please try later."
		response.Status = false
		response.StausCode = 0
		response.Error = err.Error()
		return response, nil
	}

	if exist.Email != "" {
		response.Message = "User is already exists. Please try diffrent email or login"
		response.Status = false
		response.StausCode = 2
		return response, nil
	}

	user.ID = GetNextSequenceWitS(_sn, "uid", 5000000)

	if dberr := c.Insert(user); dberr != nil {
		response.Message = "Error while user registration. Please try later."
		response.Status = false
		response.StausCode = 0
		response.Error = dberr.Error()
		return response, nil
	}

	response.Message = "Account registered successfully."
	response.Status = true
	response.StausCode = 1
	response.Error = ""
	return response, nil
}
