package golangfunctions

import (
	"fmt"
	"net/http"
)

//OlaMundo funcao de exemplo
func OlaMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ola BWG...")
}
