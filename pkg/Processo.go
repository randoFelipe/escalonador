//  O simulador deve ter como entrada as informações de cada processo como
//  PID, duração, tempo de chegada, caso tenha operação de I/O quando elas
//  devem ser executadas (em relação ao seu tempo de execução). O tempo
//  do quantum também deve ser descrito no início da simulação.
package Processo

type Processo struct{
    pid int64
    duracao int64
    tempoDeChegada int64
    io bool

	//imprime as informacoes relacionadas ao processo
    func (p *Processo) imprime() string{
        return "O processo %s tem duracao de %d e tempo de chegada %s",pid,duracao,tempoDeChegada
    }
}