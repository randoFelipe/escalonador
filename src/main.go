package main

import "pkg/Processo"



proc1 = Processo(1, 9, 10, False)
proc2 = Processo(2,10,4,False)
proc3 = Processo(3,5,0, False)
proc4 = Processo(4,7,1, False)
proc5 = Processo(5,2,17, False)

listaDeProcessos = [proc1 ,proc2 ,proc3 ,proc4 ,proc5]

def ordenaProcessos (arrayDeProcessos):
    """
    Array para ordenar processos pela ordem de chegada
    """
    aux = []
    for processo in arrayDeProcessos:
        aux.append(processo.tempoDeChegada)
    aux.sort()
    return aux


print(ordenaProcessos(listaDeProcessos))