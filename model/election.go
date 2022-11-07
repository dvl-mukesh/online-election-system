package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Election struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Location       string             `bson:"location" json:"location"`
	ElectionDate   string             `bson:"election_date" json:"election_date"`
	ResultDate     string             `bson:"result_date" json:"result_date"`
	Result         string             `bson:"result" json:"password"`
	ElectionStatus string             `bson:"election_status" json:"phone_number"`
	Candidates     []Candidate        `bson:"candidates" json:"candidates"`
}

type Candidate struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name                 string             `bson:"name" json:"name"`
	VoteCount            int64              `bson:"vote_count" json:"vote_count"`
	VoteSign             string             `bson:"vote_sign" json:"vote_sign"`
	NominationStatus     string             `bson:"nomination_status" json:"nomination_status"`
	NominationVerifiedBy primitive.ObjectID `bson:"nomination_verified_by,omitempty" json:"nomination_verified_by,omitempty"`
}
