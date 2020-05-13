package middlew

import (
	"net/http"

	"github.com/Lissen01/UndTwitter/bd"
)

/*ChequeoBD hace un checking para ver si la conexion sigue estable o no*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion Perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
