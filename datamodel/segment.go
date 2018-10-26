package datamodel

import "time"

type SegmentWrapper struct {
	TravelTime   string          `bson:"travel_tm" json:"travel_tm"`
	ToalDistance float64         `bson:"total_distance" json:"total_distance"`
	Segments     []SegmentArr    `bson:"segment" json:"segment"`
	ACCSegments  []AccSegmentArr `bson:"acc_segment" json:"acc_segment"`
	MaxSpeed     int             `bson:"mx_spd" json:"mx_spd"`
	AvgSpeed     int             `bson:"avg_spd" json:"avg_spd"`
	Date         time.Time       `bson:"date" json:"date"`
	Vhid         string          `bson:"vhid" json:"vhid"`
	AccAvail     bool            `bson:"accavl" json:"accavl"`
}

type SegmentArr struct {
	Distance    float64   `bson:"dist" json:"dist"`
	Duration    string    `bson:"dur" json:"dur"`
	StartTm     time.Time `bson:"sttm" json:"sttm"`
	EndTm       time.Time `bson:"entm" json:"entm"`
	EncodPoly   string    `bson:"encdpoly" json:"encdpoly"`
	MaxSpeed    int       `bson:"mxspd" json:"mxspd"`
	MaxSpeedLoc []float64 `bson:"mxloc" json:"mxloc"`
	MaxSpeedTM  time.Time `bson:"mxtm" json:"mxtm"`
	TrackType   string    `bson:"trktyp" json:"trktyp"`
}

type AccSegmentArr struct {
	Duration   string    `bson:"dur" json:"dur"`
	StartTm    time.Time `bson:"sttm" json:"sttm"`
	EndTm      time.Time `bson:"entm" json:"entm"`
	Point      []float64 `bson:"loc" json:"loc"`
	Address    string    `bson:"addr" json:"addr"`
	Type       string    `bson:"segtyp" json:"segtyp"`
	IsContinue bool      `bson:"isctnue" json:"isctnue"`
}

type Locdt struct {
	Time  time.Time `bson:"tm" json:"tm"`
	GpTm  string    `bson:"gpstm" json:"gpstm"`
	Loc   []float64 `bson:"l" json:"l"`
	Speed int       `bson:"spd" json:"spd"`
}

type Polocdt struct {
	Loc []float64
}
