package datamodel

import (
	"time"
)

type (
	// Vehicles represents the structure of our resource
	UserLoginLog struct {
		ID         int         `bson:"_id" json:"_id"`
		UID        int         `bson:"uid" json:"uid"`
		Email      string      `bson:"email" json:"email"`
		LoginTime  time.Time   `bson:"login_time" json:"-"`
		Source     string      `bson:"src" json:"src"` // from google, facebook
		LogoutTime time.Time   `bson:"logouttime" json:"logouttime"`
		Extra      interface{} `bson:"extra" json:"-"`
	}
)
