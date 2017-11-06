# escalonador
Esta atividade consiste na implementação de um programa que simule um
escalonador Round-Robin preemptivo conforme os estudos realizados durante
as aulas de Sistemas Operacionais.
O escalonador deve contemplar o funcionamento usual do algoritmo e
também deve possuir a funcionalidade de haver preempção devido a
operação de I/O que cada processo possa solicitar. Desta forma o
escalonador deve levar em consideração o quantum de tempo que um
processo pode utilizar a CPU, assim como quando o processo necessitar uma
operação de I/O deve ser retirado de execução na CPU. Assim todo processo
que for retirado da CPU pelo escalonador, seja porque o quantum expirou ou
por necessidade de uma operação de I/O, ele deve ser colocado no final da
Fila de Pronto (fila de espera). Caso ocorra de um novo processo chegar no
mesmo instante em que um processo que estava em execução e foi retirado
da CPU para a fila de espera, o processo em execução terá prioridade em
relação ao novo processo.
O simulador deve ter como entrada as informações de cada processo como
PID, duração, tempo de chegada, caso tenha operação de I/O quando elas
devem ser executadas (em relação ao seu tempo de execução). O tempo
do quantum também deve ser descrito no início da simulação.
Após a leitura de dados, o simulador deve apresentar o resultado de
execução dos processos, como um diagrama de Gantt, calculando o tempo
de espera de cada processo e o tempo de espera médio.
Exemplo:
quantum = 4
PID Duração Chegada Operação
I/O
P1 9 10 2, 4, 6, 8
P2 10 4 5
P3 5 0 2
P4 7 1 3, 6
P5 2 17 -
Arquivo de entrada
A Figura 1 e a Figura 2 ilustram exemplos de arquivos de entrada válido:
1. A primeira linha mostra o nome do processo (PID)
2. A segunda linha a duração do processo;
3. A terceira linha o instante de chegada do processo;
4. A quarta linha indica se o processo realiza operações de I/O;
5. Caso o processo tenha operações de I/O, a quinta linha deve indicar
em que instante essas operações ocorrem. Caso contrário a linha pode
ficar vazia.
Figura 1 – Exemplo de Arquivo de Entrada com operação de I/O
Figura 2 – Exemplo de Arquivo de Entrada sem operação de I/O
Saída de dados
A saída de dados pode ser realizada imprimindo o resultado no terminal (saída
padrão) ou em um arquivo de saída. O exemplo a seguir ilustra os dados que
devem ser apresentados:
P1
9
10
SIM
2 4 6 8
P5
2
17
NAO
***************************************
******* Escalonador Round-Robin *******
***************************************
Tempo 0 chegada de processo P3
FILA: Nao ha processos na fila
CPU: P3(5)
Tempo 1 chegada de processo P4
FILA: P4(7)
CPU: P3(4)
Tempo 2 operacao de I/O P3
FILA: P3(3)
CPU: P4(7)
Tempo 3
FILA: P3(3)
CPU: P4(6)
Tempo 4 chegada de processo P2
FILA: P3(3) P2(10)
CPU: P4(5)
Tempo 5 operacao de I/O P4
FILA: P2(10) P4(4)
CPU: P3(3)
Tempo 6
FILA: P2(10) P4(4)
CPU: P3(2)
Tempo 7
FILA: P2(10) P4(4)
CPU: P3(1)
Tempo 8 fim de processo P3
FILA: P4(4)
CPU: P2(10)
Tempo 9
FILA: P4(4)
CPU: P2(9)
Tempo 10 chegada de processo P1
FILA: P4(4) P1(9)
CPU: P2(8)
Tempo 11
FILA: P4(4) P1(9)
CPU: P2(7)
Tempo 12 fim de quantum P2
FILA: P1(9) P2(6)
CPU: P4(4)
Tempo 13
FILA: P1(9) P2(6)
CPU: P4(3)
Tempo 14
FILA: P1(9) P2(6)
CPU: P4(2)
Tempo 15 operacao de I/O P4
FILA: P2(6) P4(1)
CPU: P1(9)
Tempo 16
FILA: P2(6) P4(1)
CPU: P1(8)
Tempo 17 chegada de processo P5
operacao de I/O P1
FILA: P4(1) P5(2) P1(7)
CPU: P2(6)
Tempo 18 operacao de I/O P2
FILA: P5(2) P1(7) P2(5)
CPU: P4(1)
Tempo 19 fim de processo P4
FILA: P1(7) P2(5)
CPU: P5(2)
Tempo 20
FILA: P1(7) P2(5)
CPU: P5(1)
Tempo 21 fim de processo P5
FILA: P2(5)
CPU: P1(7)
Tempo 22
FILA: P2(5)
CPU: P1(6)
Tempo 23 operacao de I/O P1
FILA: P1(5)
CPU: P2(5)
Tempo 24
FILA: P1(5)
CPU: P2(4)
Tempo 25
FILA: P1(5)
CPU: P2(3)
Tempo 26
FILA: P1(5)
CPU: P2(2)
Tempo 27 fim de quantum P2
FILA: P2(1)
CPU: P1(5)
Tempo 28
FILA: P2(1)
CPU: P1(4)
Tempo 29 operacao de I/O P1
FILA: P1(3)
CPU: P2(1)
Tempo 30 fim de processo P2
FILA: Nao ha processos na fila
CPU: P1(3)
Tempo 31
FILA: Nao ha processos na fila
CPU: P1(2)
Tempo 32 operacao de I/O P1
FILA: Nao ha processos na fila
CPU: P1(1)
Tempo 33 fim de processo P1
FILA: Nao ha processos na fila
***************************************
Encerrando simulacao de escalonamento
***************************************
