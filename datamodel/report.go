package datamodel

type ReportParams struct {
	Reporttyp string      `bson:"reporttyp" json:"reporttyp"`
	Params    interface{} `bson:"params" json:"params"`
}
