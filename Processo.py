class Processo:
    """
    O simulador deve ter como entrada as informações de cada processo como
    PID, duração, tempo de chegada, caso tenha operação de I/O quando elas
    devem ser executadas (em relação ao seu tempo de execução). O tempo
    do quantum também deve ser descrito no início da simulação.
    """
    def __init__(self, pid, duracao, tempoDeChegada, io):
        self.pid = pid
        self.duracao = duracao
        self.tempoDeChegada = tempoDeChegada
        self.io = io
    
    def imprime(self):
        """
        imprime as informacoes relacionadas ao processo
        """
        print("O processo {} tem duracao de {} e tempo de chegada {}".format(self.pid,self.duracao,self.tempoDeChegada))

//  O simulador deve ter como entrada as informações de cada processo como
//  PID, duração, tempo de chegada, caso tenha operação de I/O quando elas
//  devem ser executadas (em relação ao seu tempo de execução). O tempo
//  do quantum também deve ser descrito no início da simulação.
type Processo struct{
    pid int64
    duracao int64
    tempoDeChegada int64
    io bool

    func (p *Processo) imprime() string{
        return "O processo %s tem duracao de %d e tempo de chegada %s",pid,duracao,tempoDeChegada
    }
}