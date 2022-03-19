package bd

import (
	"context"
	"github.com/luguerrero-meli/twittorLuis/models"
	"time"
)

func BorroRelacion(r models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	bd := MongoConnec.Database("twittor")
	col:= bd.Collection("relation")

	_, err := col.DeleteOne(ctx, r)
	if err != nil {
		return false, err
	}
	return true, nil
}
