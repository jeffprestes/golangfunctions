package testfunctions

import (
	"fmt"
	"net/http"
)

//SouUmTeste retornar o IP do usuario
func SouUmTeste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.UserAgent()+" seu IP é "+r.RemoteAddr+" ?")
}
