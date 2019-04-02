package testfunctions

import (
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("Coloque aqui algo que você quer inicializar antes da execução da Cloud Function")
}

//SouUmTeste retornar o IP do usuario
func SouUmTeste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.UserAgent()+" seu IP é "+r.RemoteAddr+" ?")
}
