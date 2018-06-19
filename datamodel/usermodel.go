package datamodel

import (
	"time"
)

type (
	// Vehicles represents the structure of our resource
	User struct {
		ID           int       `bson:"id" json:"id"`
		Email        string    `bson:"email" json:"email"`
		FirstName    string    `bson:"first_name" json:"first_name"`
		LastName     string    `bson:"last_name" json:"last_name"`
		DisplayName  string    `bson:"display_name" json:"display_name"`
		Created      time.Time `bson:"create_time" json:"-"`
		Updated      time.Time `bson:"update_time" json:"-"`
		Mobile       string    `bson:"mob" json:"mob"`
		ProfilePic   string    `bson:"prof_pic" json:"prof_pic"`
		Gender       string    `bson:"gender" json:"gender"`
		DOB          string    `bson:"dob" json:"dob"`
		Source       string    `bson:"src" json:"src"`   // from google, facebook
		VehiclesIMEI []string  `bson:"vhsimei" json:"-"` //n list of imei numbers belongs to
		Role         []string  `bson:"role" json:"-"`
		Active       bool      `bson:"active" json:"active"`
		Session      int       `bson:"session" json:"session"`
		LastLogin    time.Time `bson:"lstlogin" json:"lstlogin"`
	}
)

type (
	UserSelect struct {
		ID          int       `bson:"id" json:"id"`
		Email       string    `bson:"email" json:"email"`
		FirstName   string    `bson:"first_name" json:"first_name"`
		LastName    string    `bson:"last_name" json:"last_name"`
		DisplayName string    `bson:"display_name" json:"display_name"`
		Mobile      string    `bson:"mob" json:"mob"`
		ProfilePic  string    `bson:"prof_pic" json:"prof_pic"`
		Gender      string    `bson:"gender" json:"gender"`
		DOB         string    `bson:"dob" json:"dob"`
		Source      string    `bson:"src" json:"src"` // from google, facebook
		Active      bool      `bson:"active" json:"active"`
		Session     int       `bson:"session" json:"session"`
		LastLogin   time.Time `bson:"lstlogin" json:"lstlogin"`
	}
)
