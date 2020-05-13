package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de Conexion a la BD*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://Lissen:ngh8lmrZSROpmUUF@undtwitter-kgrva.mongodb.net/test?retryWrites=true&w=majority")

/*ConectarBD Funcion que permite la conexion a la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	/*Se hace un Ping para verificar la conexion a la BD*/
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa a la BD")
	return client
}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
