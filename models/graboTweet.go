package models

import "time"

/*GraboTweet es el formato o estructura que tendra el Twwet en la base de datos Mongo*/
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
