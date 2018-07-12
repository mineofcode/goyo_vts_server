package models

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/sender"
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

	if user.Mobile == "" {
		response.Message = "Mobile no is required"
		response.Status = false
		response.StausCode = 0
		return response, nil
	}
	c := col(_sn, db.ColUsers)

	var exist datamodel.User
	err = c.Find(bson.M{"mob": user.Mobile}).One(&exist)
	if err == nil && exist.Mobile != "" {
		response.Message = "Mobile no is already exists."
		response.Status = false
		response.StausCode = 0
		response.Error = err.Error()
		return response, nil
	}
	exist = datamodel.User{}
	user.Created = time.Now()
	user.Updated = time.Now()

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

	user.ID = GetNextSequenceWitS(_sn, SEQUserID, 5000000)
	user.OTP = utils.GenerateOtp()
	expireOtpInMin := 5
	user.Active = false
	user.OTPExp = time.Now().Add(5 * time.Minute)
	if dberr := c.Insert(user); dberr != nil {
		response.Message = "Error while user registration. Please try later."
		response.Status = false
		response.StausCode = 0
		response.Error = dberr.Error()
		return response, nil
	}
	//Send Otp To Mobile
	sendOtp(user.Mobile, user.OTP, expireOtpInMin)
	response.Message = "Account registered successfully."
	response.Status = true
	response.StausCode = 1
	response.Error = ""
	return response, nil
}

func ResendOtp(user datamodel.User) (response utils.Response) {

	_sn := getDBSession().Copy()
	defer _sn.Close()

	/*Pre checking*/
	response = utils.Response{}

	if user.Mobile == "" {
		response.Message = "Mobile is required"
		response.Status = false
		response.StausCode = 0
		return response
	}

	c := col(_sn, db.ColUsers)
	var exist datamodel.User
	query := c.Find(bson.M{"mob": user.Mobile})
	err := query.One(&exist)
	if err != nil && err.Error() != "not found" {
		response.Message = "Error while user registration. Please try later."
		response.Status = false
		response.StausCode = 0
		response.Error = err.Error()
		return response
	}

	otp := utils.GenerateOtp()
	expireOtpInMin := 5
	otpexp := time.Now().Add(5 * time.Minute)

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"otp": otp, "otpexp": otpexp}},
		Upsert:    false,
		ReturnNew: true,
	}
	var doc = datamodel.User{}
	_, err = query.Apply(change, &doc)
	if err != nil {
		response.Message = "Error while generate otp. Please try later."
		response.Status = false
		response.StausCode = 0
		response.Error = err.Error()
		return response
	}

	sendOtp(user.Mobile, otp, expireOtpInMin)
	response.Message = "OTP send successfully."
	response.Status = true
	response.StausCode = 1
	response.Error = ""
	return response
}

//RegisterUser for registering user
func VerifyOtp(user datamodel.User) (response utils.Response, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	/*Pre checking*/

	response = utils.Response{}

	if user.Mobile == "" {
		response.Message = "Mobile Number is required"
		response.Status = false
		response.StausCode = 0
		return response, nil
	}

	c := col(_sn, db.ColUsers)
	var exist datamodel.User
	query := c.Find(bson.M{"mob": user.Mobile})
	err = query.One(&exist)
	if err != nil && err.Error() == "not found" {
		response.Message = "Invalid user!"
		response.Status = false
		response.StausCode = 2
		response.Error = err.Error()
		return response, nil
	}
	if exist.Active {
		response.Message = "OTP already verified!"
		response.Status = false
		response.StausCode = 3
		return response, nil
	}
	if exist.OTP != user.OTP {
		response.Message = "Invalid OTP!"
		response.Status = false
		response.StausCode = 3
		return response, nil
	}

	if exist.OTPExp.Before(time.Now()) {
		response.Message = "OTP Expired!"
		response.Status = false
		response.StausCode = 4
		return response, nil
	}

	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"active": true, "acttime": time.Now()}},
		Upsert:    false,
		ReturnNew: true,
	}

	var doc = datamodel.User{}
	_, err = query.Apply(change, &doc)
	if err != nil {
		response.Message = "Error while user registration. Please try later."
		response.Status = false
		response.StausCode = 0
		response.Error = err.Error()
		return response, nil
	}

	response.Message = "Mobile no. verification successfully."
	response.Status = true
	response.StausCode = 1
	response.Error = ""
	return response, nil
}

func sendOtp(mobile string, otp string, expires int) {
	search := make(map[string]interface{})
	search["key"] = "sms"
	search["group"] = "settings"
	result := GetMoM(search)
	if !result.Status {
		return
	}
	data := result.Data.([]STRUCTMom)
	values := data[0].Valued.(bson.M)
	go sender.SendSMS(mobile,
		fmt.Sprintf("Your Account Verification OTP: %s Expired in %d minutes", otp, expires),
		values)
}
