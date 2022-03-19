package bd

import (
	"context"
	"github.com/luguerrero-meli/twittorLuis/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func LeoTweetSeguidores(ID string, page int) ([]*models.DevuelvoTweetSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnec.Database("twittor")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	condiciones := make([]bson.M,0)
	condiciones = append(condiciones, bson.M{"$match":bson.M{"usuarioid":ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":"tweet",
			"localField":"usuariorelacionid",
			"foreignField":"userId",
			"as":"tweet",
		},})
	condiciones = append(condiciones, bson.M{"$unwind":"$tweet"})
	condiciones = append(condiciones, bson.M{"$sort":bson.M{"date":-1}})
	condiciones = append(condiciones, bson.M{"$skip":skip})
	condiciones = append(condiciones, bson.M{"$limit":20})

	var result []*models.DevuelvoTweetSeguidores
	cursor, err := col.Aggregate(ctx,condiciones)
	if err != nil {
		return result, false
	}
	err = cursor.All(ctx,&result)
	if err != nil {
		return result,false
	}
	return result,true
}
