package main

import (
	"fmt"
)

type Processo struct {
	Pid              string
	Duracao, Chegada int8
	Io               bool
}

func escalonador(listaDeProcessos []Processo, quantum int8) {
	var processosNaFila []Processo
	var tempo int8

}

func main() {
	fmt.Println("*************\nEscalonador Round-Robin\n*************")

}
