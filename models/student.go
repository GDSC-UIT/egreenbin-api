package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID             primitive.ObjectID `bson:"_id" json:"id" `
	Code           string             `bson:"code" json:"code"`
	Name           string             `bson:"name" json:"name"`
	NumOfCorrect   int                `bson:"numOfCorrect" json:"numOfCorrect"`
	NumOfWrong     int                `bson:"numOfWrong" json:"numOfWrong"`
	ImageAvatarUrl string             `bson:"imageAvatarUrl" json:"imageAvatarUrl"`
	ParentEmail    string             `bson:"parentEmail" json:"parentEmail"`
	Note           string             `bson:"note" json:"note"`
}
