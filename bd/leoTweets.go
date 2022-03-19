package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/luguerrero-meli/twittorLuis/models"
)

/*LeoTweets permite leer los tweets de un perfil*/
func LeoTweets(ID string, page int64) ([]*models.DevuelveTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoConnec.Database("twittor")
	col := bd.Collection("tweet")

	var resultado []*models.DevuelveTweets

	condicion := bson.M{
		"userId": ID,
	}

	opciones := options.Find()
	opciones.SetSkip((page - 1) * 20)
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultado, false
	}

	for cursor.Next(context.TODO()) {
		var registro models.DevuelveTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, false
		}
		resultado = append(resultado, &registro)
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	cursor.Close(ctx)

	return resultado, true
}
