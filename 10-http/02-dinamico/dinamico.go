package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	/*
		FORMATAR DATA
		01 => dia
		02 => mes
		03 => hora
		04 => minuto
		05 => segundo
	*/
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s </h1>", s)
}

func main() {
	//"obriga" que uma determinada url chame determidado metodo
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Excutando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
