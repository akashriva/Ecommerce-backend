package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Verification Model `verifications` table
type Verification struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Email      string             `bson:"email" json:"email"`
	Status     bool               `bson:"status" json:"status"`
	Otp        string             `bson:"otp" json:"otp"`
	OtpExpire  time.Time          `bson:"otp_expire" json:"otp_expire"`
	UserId     primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
}

