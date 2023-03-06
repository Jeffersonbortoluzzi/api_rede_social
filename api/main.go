package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Função para gerar uma única vez a secret key (está no .env)
// func init() {
// 	chave := make([]byte, 64)

// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	// transformar o slice de byte em string base64 para colocar no .env
// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)

// }

func main() {
	config.Carregar()
	r := router.Gerar()

	// fmt.Println(config.SecretKey) testando para confirmar que está passando em forma de slice de bytes

	fmt.Printf("Rodando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
