package main

import (
	"fmt"
)

type Processo struct {
	Pid              string
	Duracao, Chegada int8
	Io               bool
}

func main() {
	p1 := Processo{"P1", 9, 10, false}
	p2 := Processo{"P2", 10, 4, false}
	p3 := Processo{"P3", 5, 0, false}
	p4 := Processo{"P4", 7, 1, false}
	p5 := Processo{"P5", 2, 17, false}
	listaDeProcessos := [5]Processo{p1, p2, p3, p4, p5}
	quantum := 4
	fmt.Println("*************\nEscalonador Round-Robin\n*************")

	fmt.Println("*************\nEncerrando simulação de escalonamento\n*************")

}

//função que realiza a logica do escalonador
func escalonador(listaDeProcessos []Processo) {
	tempo := 0

}

//para cada tempo de duracao de um quantum a função subtrai 1 de duração de um processo
func cpuExecutor(p Processo, quantum int) (duracao int) {
	for i := 1; i <= quantum; i++ {
		if p.Duracao > 0 {
			p.Duracao -= 1
		}

	}
	return
}

//função responsavel por realocar os objetos dentro de uma array
//ela pega o primeiro elemento da array e coloca no fim, realocando todos os outros objetos para
//sua posição -1
func reArrange(lista *[5]Processo) {
	var aux Processo
	for i := 0; i < len(lista)-1; i++ {
		if i == 0 {
			aux = lista[0]
		}
		lista[i] = lista[i+1]
	}
	lista[len(lista)-1] = aux
}
