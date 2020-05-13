package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword rutina que me permite encriptar contraseñas a tipo SHA1024*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
