package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoConnec es la variable de coneccion*/
var MongoConnec = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://luguerrero:BKPCuLMtqDeBggoM@cluster0.kztyh.mongodb.net/twittor?retryWrites=true&w=majority")

/*ConectarBD es la funcion de conectar a base de datos*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la DB")
	return client
}

/*ChequeoConnection es la funcion de validacion de coneccion*/
func ChequeoConnection() int {
	err := MongoConnec.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
