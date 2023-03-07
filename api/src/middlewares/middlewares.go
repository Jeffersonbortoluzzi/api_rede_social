package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// Logger escreve dados da requisição dentro do terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Auntenticar verifíca se o usuario que fez a autenticação está autenticado
func Auntenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auntenticando")
		next(w, r)
	}
}
