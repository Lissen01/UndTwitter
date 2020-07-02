package models

/*RespuestaLogin tiene el token que se devuelve al login*/
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
