package bd

import (
	"context"
	"time"

	"github.com/Lissen01/UndTwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibe un mail de parametro y chequea si el mismo esta en la DB*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("UndTwitter")

	col := db.Collection("usuarios")

	condicion := bson.M{"email": email} /*Formatea en formato bson*/

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
