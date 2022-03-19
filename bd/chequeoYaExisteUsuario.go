package bd

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/luguerrero-meli/twittorLuis/models"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	fmt.Println("Email entrante: " + email)
	db := MongoConnec.Database("twittor")
	col := db.Collection("users")

	condicion := bson.M{"email": email}
	fmt.Println(condicion)
	var resultado models.Usuario
	fmt.Println(ctx)

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
