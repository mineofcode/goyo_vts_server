package datamodel

type (
	// Vehicles represents the structure of our resource
	Counters struct {
		ID  string `bson:"_id"`
		Seq int    `bson:"value"`
	}
)
